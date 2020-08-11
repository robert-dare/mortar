// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package mock_trace is a generated GoMock package.
package mock_trace

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	opentracing "github.com/opentracing/opentracing-go"
	reflect "reflect"
)

// MockOpenTracer is a mock of OpenTracer interface.
type MockOpenTracer struct {
	ctrl     *gomock.Controller
	recorder *MockOpenTracerMockRecorder
}

// MockOpenTracerMockRecorder is the mock recorder for MockOpenTracer.
type MockOpenTracerMockRecorder struct {
	mock *MockOpenTracer
}

// NewMockOpenTracer creates a new mock instance.
func NewMockOpenTracer(ctrl *gomock.Controller) *MockOpenTracer {
	mock := &MockOpenTracer{ctrl: ctrl}
	mock.recorder = &MockOpenTracerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenTracer) EXPECT() *MockOpenTracerMockRecorder {
	return m.recorder
}

// Connect mocks base method.
func (m *MockOpenTracer) Connect(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockOpenTracerMockRecorder) Connect(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockOpenTracer)(nil).Connect), ctx)
}

// Tracer mocks base method.
func (m *MockOpenTracer) Tracer() opentracing.Tracer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tracer")
	ret0, _ := ret[0].(opentracing.Tracer)
	return ret0
}

// Tracer indicates an expected call of Tracer.
func (mr *MockOpenTracerMockRecorder) Tracer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tracer", reflect.TypeOf((*MockOpenTracer)(nil).Tracer))
}

// Close mocks base method.
func (m *MockOpenTracer) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockOpenTracerMockRecorder) Close(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockOpenTracer)(nil).Close), ctx)
}