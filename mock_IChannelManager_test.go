// Code generated by mockery v1.0.0. DO NOT EDIT.

package amqpwrapper

import (
	amqp "github.com/streadway/amqp"
	mock "github.com/stretchr/testify/mock"
)

// MockIChannelManager is an autogenerated mock type for the IChannelManager type
type MockIChannelManager struct {
	mock.Mock
}

// Consume provides a mock function with given fields: queue, consumer, autoAck, exclusive, noLocal, noWait, args
func (_m *MockIChannelManager) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	ret := _m.Called(queue, consumer, autoAck, exclusive, noLocal, noWait, args)

	var r0 <-chan amqp.Delivery
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, bool, bool, amqp.Table) <-chan amqp.Delivery); ok {
		r0 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan amqp.Delivery)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, bool, bool, bool, bool, amqp.Table) error); ok {
		r1 = rf(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsClosed provides a mock function with given fields:
func (_m *MockIChannelManager) IsClosed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Publish provides a mock function with given fields: exchange, key, mandatory, immediate, msg
func (_m *MockIChannelManager) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {
	ret := _m.Called(exchange, key, mandatory, immediate, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, amqp.Publishing) error); ok {
		r0 = rf(exchange, key, mandatory, immediate, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
