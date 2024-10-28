// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	domain "go-redpanda-streaming/domain"

	mock "github.com/stretchr/testify/mock"
)

// StreamUsecase is an autogenerated mock type for the StreamUsecase type
type StreamUsecase struct {
	mock.Mock
}

// GetResults provides a mock function with given fields: streamID
func (_m *StreamUsecase) GetResults(streamID string) ([]domain.Message, error) {
	ret := _m.Called(streamID)

	if len(ret) == 0 {
		panic("no return value specified for GetResults")
	}

	var r0 []domain.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]domain.Message, error)); ok {
		return rf(streamID)
	}
	if rf, ok := ret.Get(0).(func(string) []domain.Message); ok {
		r0 = rf(streamID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(streamID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendData provides a mock function with given fields: streamID, data
func (_m *StreamUsecase) SendData(streamID string, data domain.StreamData) error {
	ret := _m.Called(streamID, data)

	if len(ret) == 0 {
		panic("no return value specified for SendData")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, domain.StreamData) error); ok {
		r0 = rf(streamID, data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartStream provides a mock function with given fields: streamID
func (_m *StreamUsecase) StartStream(streamID string) error {
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

// NewStreamUsecase creates a new instance of StreamUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStreamUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *StreamUsecase {
	mock := &StreamUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}