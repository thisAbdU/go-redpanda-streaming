// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	domain "go-redpanda-streaming/domain"

	mock "github.com/stretchr/testify/mock"
)

// StreamRepository is an autogenerated mock type for the StreamRepository type
type StreamRepository struct {
	mock.Mock
}

// ReceiveMessages provides a mock function with given fields: streamID
func (_m *StreamRepository) ReceiveMessages(streamID string) (<-chan domain.Message, error) {
	ret := _m.Called(streamID)

	if len(ret) == 0 {
		panic("no return value specified for ReceiveMessages")
	}

	var r0 <-chan domain.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (<-chan domain.Message, error)); ok {
		return rf(streamID)
	}
	if rf, ok := ret.Get(0).(func(string) <-chan domain.Message); ok {
		r0 = rf(streamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan domain.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(streamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: streamID, message
func (_m *StreamRepository) SendMessage(streamID string, message domain.Message) error {
	ret := _m.Called(streamID, message)

	if len(ret) == 0 {
		panic("no return value specified for SendMessage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.Message) error); ok {
		r0 = rf(streamID, message)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartStream provides a mock function with given fields: streamID
func (_m *StreamRepository) StartStream(streamID string) error {
	ret := _m.Called(streamID)

	if len(ret) == 0 {
		panic("no return value specified for StartStream")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(streamID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStreamRepository creates a new instance of StreamRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStreamRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *StreamRepository {
	mock := &StreamRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
