// Code generated by MockGen. DO NOT EDIT.
// Source: cs3219-project-ay2223s1-g33/session-service/blacklist (interfaces: RedisBlacklistClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	blacklist "cs3219-project-ay2223s1-g33/session-service/blacklist"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRedisBlacklistClient is a mock of RedisBlacklistClient interface.
type MockRedisBlacklistClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedisBlacklistClientMockRecorder
}

// MockRedisBlacklistClientMockRecorder is the mock recorder for MockRedisBlacklistClient.
type MockRedisBlacklistClientMockRecorder struct {
	mock *MockRedisBlacklistClient
}

// NewMockRedisBlacklistClient creates a new mock instance.
func NewMockRedisBlacklistClient(ctrl *gomock.Controller) *MockRedisBlacklistClient {
	mock := &MockRedisBlacklistClient{ctrl: ctrl}
	mock.recorder = &MockRedisBlacklistClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedisBlacklistClient) EXPECT() *MockRedisBlacklistClientMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockRedisBlacklistClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRedisBlacklistClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRedisBlacklistClient)(nil).Close))
}

// Connect mocks base method.
func (m *MockRedisBlacklistClient) Connect() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect.
func (mr *MockRedisBlacklistClientMockRecorder) Connect() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockRedisBlacklistClient)(nil).Connect))
}

// GetRefreshBlacklist mocks base method.
func (m *MockRedisBlacklistClient) GetRefreshBlacklist() blacklist.TokenBlacklist {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRefreshBlacklist")
	ret0, _ := ret[0].(blacklist.TokenBlacklist)
	return ret0
}

// GetRefreshBlacklist indicates an expected call of GetRefreshBlacklist.
func (mr *MockRedisBlacklistClientMockRecorder) GetRefreshBlacklist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRefreshBlacklist", reflect.TypeOf((*MockRedisBlacklistClient)(nil).GetRefreshBlacklist))
}

// GetSessionBlacklist mocks base method.
func (m *MockRedisBlacklistClient) GetSessionBlacklist() blacklist.TokenBlacklist {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSessionBlacklist")
	ret0, _ := ret[0].(blacklist.TokenBlacklist)
	return ret0
}

// GetSessionBlacklist indicates an expected call of GetSessionBlacklist.
func (mr *MockRedisBlacklistClientMockRecorder) GetSessionBlacklist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSessionBlacklist", reflect.TypeOf((*MockRedisBlacklistClient)(nil).GetSessionBlacklist))
}
