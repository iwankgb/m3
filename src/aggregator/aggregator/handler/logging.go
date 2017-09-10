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

package handler

import (
	"github.com/m3db/m3aggregator/aggregator"
	"github.com/m3db/m3metrics/metric/aggregated"
	"github.com/m3db/m3metrics/policy"
	"github.com/m3db/m3x/instrument"
	"github.com/m3db/m3x/log"
)

func logMetricAndPolicy(
	logger xlog.Logger,
	metric aggregated.Metric,
	sp policy.StoragePolicy,
) error {
	logger.WithFields(
		xlog.NewLogField("metric", metric.String()),
		xlog.NewLogField("policy", sp.String()),
	).Info("aggregated metric")
	return nil
}

func loggingHandler(logger xlog.Logger) HandleFunc {
	return func(metric aggregated.Metric, sp policy.StoragePolicy) error {
		return logMetricAndPolicy(logger, metric, sp)
	}
}

// NewLoggingHandler creates a new logging handler.
func NewLoggingHandler(instrumentOpts instrument.Options) aggregator.Handler {
	handler := loggingHandler(instrumentOpts.Logger())
	return NewDecodingHandler(handler)
}