// +build integration

package amqpwrapper

import amqp "github.com/streadway/amqp"
import mock "github.com/stretchr/testify/mock"

// MockChannelManager is an autogenerated mock type for the MockChannelManager type
type MockChannelManager struct {
	mock.Mock
}

// Consume provides a mock function with given fields: queue, consumer, autoAck, exclusive, noLocal, noWait, args
func (_m *MockChannelManager) Consume(queue string, consumer string, autoAck bool, exclusive bool, noLocal bool, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
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
func (_m *MockChannelManager) IsClosed() bool {
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
func (_m *MockChannelManager) Publish(exchange string, key string, mandatory bool, immediate bool, msg amqp.Publishing) error {
	ret := _m.Called(exchange, key, mandatory, immediate, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, bool, bool, amqp.Publishing) error); ok {
		r0 = rf(exchange, key, mandatory, immediate, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockConnectionManager is an autogenerated mock type for the MockConnectionManager type
type MockConnectionManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *MockConnectionManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateChannel provides a mock function with given fields: typeChan
func (_m *MockConnectionManager) CreateChannel(typeChan uint64) (*amqp.Channel, error) {
	ret := _m.Called(typeChan)

	var r0 *amqp.Channel
	if rf, ok := ret.Get(0).(func(uint64) *amqp.Channel); ok {
		r0 = rf(typeChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(typeChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetChannel provides a mock function with given fields: key, typeChan
func (_m *MockConnectionManager) GetChannel(key string, typeChan uint64) (*amqp.Channel, error) {
	ret := _m.Called(key, typeChan)

	var r0 *amqp.Channel
	if rf, ok := ret.Get(0).(func(string, uint64) *amqp.Channel); ok {
		r0 = rf(key, typeChan)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*amqp.Channel)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint64) error); ok {
		r1 = rf(key, typeChan)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitChannel provides a mock function with given fields: fn, args
func (_m *MockConnectionManager) InitChannel(fn InitializeChannel, args InitArgs) error {
	ret := _m.Called(fn, args)

	var r0 error
	if rf, ok := ret.Get(0).(func(InitializeChannel, InitArgs) error); ok {
		r0 = rf(fn, args)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IsClosed provides a mock function with given fields:
func (_m *MockConnectionManager) IsClosed() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// isNotValidKey provides a mock function with given fields: key
func (_m *MockConnectionManager) isNotValidKey(key string) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// isNotValidTypeChan provides a mock function with given fields: typeChan
func (_m *MockConnectionManager) isNotValidTypeChan(typeChan uint64) bool {
	ret := _m.Called(typeChan)

	var r0 bool
	if rf, ok := ret.Get(0).(func(uint64) bool); ok {
		r0 = rf(typeChan)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
