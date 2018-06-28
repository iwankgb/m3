// Copyright (c) 2018 Uber Technologies, Inc.
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

// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/m3db/m3aggregator/aggregator/flush.go

package aggregator

import (
	"time"

	"github.com/golang/mock/gomock"
)

// Mock of flushingMetricList interface
type MockflushingMetricList struct {
	ctrl     *gomock.Controller
	recorder *_MockflushingMetricListRecorder
}

// Recorder for MockflushingMetricList (not exported)
type _MockflushingMetricListRecorder struct {
	mock *MockflushingMetricList
}

func NewMockflushingMetricList(ctrl *gomock.Controller) *MockflushingMetricList {
	mock := &MockflushingMetricList{ctrl: ctrl}
	mock.recorder = &_MockflushingMetricListRecorder{mock}
	return mock
}

func (_m *MockflushingMetricList) EXPECT() *_MockflushingMetricListRecorder {
	return _m.recorder
}

func (_m *MockflushingMetricList) Shard() uint32 {
	ret := _m.ctrl.Call(_m, "Shard")
	ret0, _ := ret[0].(uint32)
	return ret0
}

func (_mr *_MockflushingMetricListRecorder) Shard() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shard")
}

func (_m *MockflushingMetricList) ID() metricListID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(metricListID)
	return ret0
}

func (_mr *_MockflushingMetricListRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockflushingMetricList) FlushInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "FlushInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockflushingMetricListRecorder) FlushInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushInterval")
}

func (_m *MockflushingMetricList) LastFlushedNanos() int64 {
	ret := _m.ctrl.Call(_m, "LastFlushedNanos")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockflushingMetricListRecorder) LastFlushedNanos() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastFlushedNanos")
}

func (_m *MockflushingMetricList) Flush(req flushRequest) {
	_m.ctrl.Call(_m, "Flush", req)
}

func (_mr *_MockflushingMetricListRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

func (_m *MockflushingMetricList) DiscardBefore(beforeNanos int64) {
	_m.ctrl.Call(_m, "DiscardBefore", beforeNanos)
}

func (_mr *_MockflushingMetricListRecorder) DiscardBefore(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscardBefore", arg0)
}

// Mock of fixedOffsetFlushingMetricList interface
type MockfixedOffsetFlushingMetricList struct {
	ctrl     *gomock.Controller
	recorder *_MockfixedOffsetFlushingMetricListRecorder
}

// Recorder for MockfixedOffsetFlushingMetricList (not exported)
type _MockfixedOffsetFlushingMetricListRecorder struct {
	mock *MockfixedOffsetFlushingMetricList
}

func NewMockfixedOffsetFlushingMetricList(ctrl *gomock.Controller) *MockfixedOffsetFlushingMetricList {
	mock := &MockfixedOffsetFlushingMetricList{ctrl: ctrl}
	mock.recorder = &_MockfixedOffsetFlushingMetricListRecorder{mock}
	return mock
}

func (_m *MockfixedOffsetFlushingMetricList) EXPECT() *_MockfixedOffsetFlushingMetricListRecorder {
	return _m.recorder
}

func (_m *MockfixedOffsetFlushingMetricList) Shard() uint32 {
	ret := _m.ctrl.Call(_m, "Shard")
	ret0, _ := ret[0].(uint32)
	return ret0
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) Shard() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Shard")
}

func (_m *MockfixedOffsetFlushingMetricList) ID() metricListID {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(metricListID)
	return ret0
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

func (_m *MockfixedOffsetFlushingMetricList) FlushInterval() time.Duration {
	ret := _m.ctrl.Call(_m, "FlushInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) FlushInterval() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushInterval")
}

func (_m *MockfixedOffsetFlushingMetricList) LastFlushedNanos() int64 {
	ret := _m.ctrl.Call(_m, "LastFlushedNanos")
	ret0, _ := ret[0].(int64)
	return ret0
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) LastFlushedNanos() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LastFlushedNanos")
}

func (_m *MockfixedOffsetFlushingMetricList) Flush(req flushRequest) {
	_m.ctrl.Call(_m, "Flush", req)
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) Flush(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Flush", arg0)
}

func (_m *MockfixedOffsetFlushingMetricList) DiscardBefore(beforeNanos int64) {
	_m.ctrl.Call(_m, "DiscardBefore", beforeNanos)
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) DiscardBefore(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DiscardBefore", arg0)
}

func (_m *MockfixedOffsetFlushingMetricList) FlushOffset() time.Duration {
	ret := _m.ctrl.Call(_m, "FlushOffset")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

func (_mr *_MockfixedOffsetFlushingMetricListRecorder) FlushOffset() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "FlushOffset")
}