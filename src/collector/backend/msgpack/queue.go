// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package msgpack

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/m3db/m3cluster/placement"
	"github.com/m3db/m3metrics/protocol/msgpack"
	"github.com/m3db/m3x/log"

	"github.com/uber-go/tally"
)

const (
	defaultQueueSizeNumBuckets = 6
)

var (
	errInstanceQueueClosed = errors.New("instance queue is closed")
	errWriterQueueFull     = errors.New("writer queue is full")
)

// DropType determines which metrics should be dropped when the queue is full.
type DropType int

const (
	// DropOldest signifies that the oldest metrics in the queue should be dropped.
	DropOldest DropType = iota

	// DropCurrent signifies that the current metrics in the queue should be dropped.
	DropCurrent
)

var (
	validDropTypes = []DropType{
		DropOldest,
		DropCurrent,
	}
)

func (t DropType) String() string {
	switch t {
	case DropOldest:
		return "oldest"
	case DropCurrent:
		return "current"
	}
	return "unknown"
}

// UnmarshalYAML unmarshals a DropType into a valid type from string.
func (t *DropType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	if err := unmarshal(&str); err != nil {
		return err
	}
	if str == "" {
		*t = defaultDropType
		return nil
	}
	strs := make([]string, len(validDropTypes))
	for _, valid := range validDropTypes {
		if str == valid.String() {
			*t = valid
			return nil
		}
		strs = append(strs, "'"+valid.String()+"'")
	}
	return fmt.Errorf(
		"invalid DropType '%s' valid types are: %s", str, strings.Join(strs, ", "),
	)
}

// instanceQueue processes write requests for given instance.
type instanceQueue interface {
	// Enqueue enqueues a data buffer.
	Enqueue(buf msgpack.Buffer) error

	// Close closes the queue.
	Close() error
}

type writeFn func([]byte) error

type queue struct {
	sync.RWMutex

	log      log.Logger
	metrics  queueMetrics
	dropType DropType
	instance placement.Instance
	conn     *connection
	bufCh    chan msgpack.Buffer
	doneCh   chan struct{}
	closed   bool

	writeFn writeFn
}

func newInstanceQueue(instance placement.Instance, opts ServerOptions) instanceQueue {
	var (
		conn      = newConnection(instance.Endpoint(), opts.ConnectionOptions())
		iOpts     = opts.InstrumentOptions()
		queueSize = opts.InstanceQueueSize()
	)
	q := &queue{
		dropType: opts.QueueDropType(),
		log:      iOpts.Logger(),
		metrics:  newQueueMetrics(iOpts.MetricsScope(), queueSize),
		instance: instance,
		conn:     conn,
		bufCh:    make(chan msgpack.Buffer, queueSize),
		doneCh:   make(chan struct{}),
	}
	q.writeFn = q.conn.Write

	go q.drain()
	go q.reportQueueSize(iOpts.ReportInterval())
	return q
}

func (q *queue) Enqueue(buf msgpack.Buffer) error {
	q.RLock()
	if q.closed {
		q.RUnlock()
		q.metrics.enqueueClosedErrors.Inc(1)
		return errInstanceQueueClosed
	}
	for {
		// NB(xichen): the buffer already batches multiple metric buf points
		// to maximize per packet utilization so there should be no need to perform
		// additional batching here.
		select {
		case q.bufCh <- buf:
			q.RUnlock()
			q.metrics.enqueueSuccesses.Inc(1)
			return nil

		default:
			if q.dropType == DropCurrent {
				q.RUnlock()

				// Close the buffer so it's resources are freed.
				buf.Close()
				q.metrics.enqueueCurrentDropped.Inc(1)
				return errWriterQueueFull
			}
		}

		select {
		case buf := <-q.bufCh:
			// Close the buffer so it's resources are freed.
			buf.Close()
			q.metrics.enqueueOldestDropped.Inc(1)
		default:
		}
	}
}

func (q *queue) Close() error {
	q.Lock()
	defer q.Unlock()

	if q.closed {
		return errInstanceQueueClosed
	}
	q.closed = true
	close(q.doneCh)
	close(q.bufCh)
	return nil
}

func (q *queue) drain() {
	for buf := range q.bufCh {
		if err := q.writeFn(buf.Bytes()); err != nil {
			q.log.WithFields(
				log.NewField("instance", q.instance.Endpoint()),
				log.NewErrField(err),
			).Error("write data error")
			q.metrics.connWriteErrors.Inc(1)
		}
		buf.Close()
	}
	q.conn.Close()
}

func (q *queue) reportQueueSize(reportInterval time.Duration) {
	ticker := time.NewTicker(reportInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			q.metrics.queueLen.RecordValue(float64(len(q.bufCh)))
		case <-q.doneCh:
			return
		}
	}
}

type queueMetrics struct {
	queueLen              tally.Histogram
	enqueueSuccesses      tally.Counter
	enqueueOldestDropped  tally.Counter
	enqueueCurrentDropped tally.Counter
	enqueueClosedErrors   tally.Counter
	connWriteErrors       tally.Counter
}

func newQueueMetrics(s tally.Scope, queueSize int) queueMetrics {
	numBuckets := defaultQueueSizeNumBuckets
	if queueSize < numBuckets {
		numBuckets = queueSize
	}
	buckets := tally.MustMakeLinearValueBuckets(0, float64(queueSize/numBuckets), numBuckets)
	enqueueScope := s.Tagged(map[string]string{"action": "enqueue"})
	return queueMetrics{
		queueLen:         s.Histogram("queue-length", buckets),
		enqueueSuccesses: enqueueScope.Counter("successes"),
		enqueueOldestDropped: enqueueScope.Tagged(map[string]string{"drop-type": "oldest"}).
			Counter("dropped"),
		enqueueCurrentDropped: enqueueScope.Tagged(map[string]string{"drop-type": "current"}).
			Counter("dropped"),
		enqueueClosedErrors: enqueueScope.Tagged(map[string]string{"error-type": "queue-closed"}).
			Counter("errors"),
		connWriteErrors: s.Tagged(map[string]string{"error-type": "conn-write", "action": "drain"}).
			Counter("errors"),
	}
}