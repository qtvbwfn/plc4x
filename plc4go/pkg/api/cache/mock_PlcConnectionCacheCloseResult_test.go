/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.42.2. DO NOT EDIT.

package cache

import mock "github.com/stretchr/testify/mock"

// MockPlcConnectionCacheCloseResult is an autogenerated mock type for the PlcConnectionCacheCloseResult type
type MockPlcConnectionCacheCloseResult struct {
	mock.Mock
}

type MockPlcConnectionCacheCloseResult_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcConnectionCacheCloseResult) EXPECT() *MockPlcConnectionCacheCloseResult_Expecter {
	return &MockPlcConnectionCacheCloseResult_Expecter{mock: &_m.Mock}
}

// GetConnectionCache provides a mock function with given fields:
func (_m *MockPlcConnectionCacheCloseResult) GetConnectionCache() PlcConnectionCache {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetConnectionCache")
	}

	var r0 PlcConnectionCache
	if rf, ok := ret.Get(0).(func() PlcConnectionCache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcConnectionCache)
		}
	}

	return r0
}

// MockPlcConnectionCacheCloseResult_GetConnectionCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetConnectionCache'
type MockPlcConnectionCacheCloseResult_GetConnectionCache_Call struct {
	*mock.Call
}

// GetConnectionCache is a helper method to define mock.On call
func (_e *MockPlcConnectionCacheCloseResult_Expecter) GetConnectionCache() *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call {
	return &MockPlcConnectionCacheCloseResult_GetConnectionCache_Call{Call: _e.mock.On("GetConnectionCache")}
}

func (_c *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call) Run(run func()) *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call) Return(_a0 PlcConnectionCache) *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call) RunAndReturn(run func() PlcConnectionCache) *MockPlcConnectionCacheCloseResult_GetConnectionCache_Call {
	_c.Call.Return(run)
	return _c
}

// GetErr provides a mock function with given fields:
func (_m *MockPlcConnectionCacheCloseResult) GetErr() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetErr")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPlcConnectionCacheCloseResult_GetErr_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetErr'
type MockPlcConnectionCacheCloseResult_GetErr_Call struct {
	*mock.Call
}

// GetErr is a helper method to define mock.On call
func (_e *MockPlcConnectionCacheCloseResult_Expecter) GetErr() *MockPlcConnectionCacheCloseResult_GetErr_Call {
	return &MockPlcConnectionCacheCloseResult_GetErr_Call{Call: _e.mock.On("GetErr")}
}

func (_c *MockPlcConnectionCacheCloseResult_GetErr_Call) Run(run func()) *MockPlcConnectionCacheCloseResult_GetErr_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcConnectionCacheCloseResult_GetErr_Call) Return(_a0 error) *MockPlcConnectionCacheCloseResult_GetErr_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcConnectionCacheCloseResult_GetErr_Call) RunAndReturn(run func() error) *MockPlcConnectionCacheCloseResult_GetErr_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPlcConnectionCacheCloseResult creates a new instance of MockPlcConnectionCacheCloseResult. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPlcConnectionCacheCloseResult(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPlcConnectionCacheCloseResult {
	mock := &MockPlcConnectionCacheCloseResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
