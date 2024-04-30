/*
Copyright (c) Facebook, Inc. and its affiliates.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: time/ptp/sptp/client/stats.go

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"
	time "time"

	stats "github.com/facebook/time/ptp/sptp/stats"
	gomock "github.com/golang/mock/gomock"
)

// MockStatsServer is a mock of StatsServer interface.
type MockStatsServer struct {
	ctrl     *gomock.Controller
	recorder *MockStatsServerMockRecorder
}

// MockStatsServerMockRecorder is the mock recorder for MockStatsServer.
type MockStatsServerMockRecorder struct {
	mock *MockStatsServer
}

// NewMockStatsServer creates a new mock instance.
func NewMockStatsServer(ctrl *gomock.Controller) *MockStatsServer {
	mock := &MockStatsServer{ctrl: ctrl}
	mock.recorder = &MockStatsServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStatsServer) EXPECT() *MockStatsServerMockRecorder {
	return m.recorder
}

// CollectSysStats mocks base method.
func (m *MockStatsServer) CollectSysStats() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectSysStats")
	ret0, _ := ret[0].(error)
	return ret0
}

// CollectSysStats indicates an expected call of CollectSysStats.
func (mr *MockStatsServerMockRecorder) CollectSysStats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectSysStats", reflect.TypeOf((*MockStatsServer)(nil).CollectSysStats))
}

// IncRXAnnounce mocks base method.
func (m *MockStatsServer) IncRXAnnounce() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncRXAnnounce")
}

// IncRXAnnounce indicates an expected call of IncRXAnnounce.
func (mr *MockStatsServerMockRecorder) IncRXAnnounce() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncRXAnnounce", reflect.TypeOf((*MockStatsServer)(nil).IncRXAnnounce))
}

// IncRXDelayReq mocks base method.
func (m *MockStatsServer) IncRXDelayReq() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncRXDelayReq")
}

// IncRXDelayReq indicates an expected call of IncRXDelayReq.
func (mr *MockStatsServerMockRecorder) IncRXDelayReq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncRXDelayReq", reflect.TypeOf((*MockStatsServer)(nil).IncRXDelayReq))
}

// IncRXSync mocks base method.
func (m *MockStatsServer) IncRXSync() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncRXSync")
}

// IncRXSync indicates an expected call of IncRXSync.
func (mr *MockStatsServerMockRecorder) IncRXSync() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncRXSync", reflect.TypeOf((*MockStatsServer)(nil).IncRXSync))
}

// IncTXDelayReq mocks base method.
func (m *MockStatsServer) IncTXDelayReq() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncTXDelayReq")
}

// IncTXDelayReq indicates an expected call of IncTXDelayReq.
func (mr *MockStatsServerMockRecorder) IncTXDelayReq() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncTXDelayReq", reflect.TypeOf((*MockStatsServer)(nil).IncTXDelayReq))
}

// IncUnsupported mocks base method.
func (m *MockStatsServer) IncUnsupported() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IncUnsupported")
}

// IncUnsupported indicates an expected call of IncUnsupported.
func (mr *MockStatsServerMockRecorder) IncUnsupported() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncUnsupported", reflect.TypeOf((*MockStatsServer)(nil).IncUnsupported))
}

// SetGMStats mocks base method.
func (m *MockStatsServer) SetGMStats(stat *stats.Stat) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGMStats", stat)
}

// SetGMStats indicates an expected call of SetGMStats.
func (mr *MockStatsServerMockRecorder) SetGMStats(stat interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGMStats", reflect.TypeOf((*MockStatsServer)(nil).SetGMStats), stat)
}

// SetGmsAvailable mocks base method.
func (m *MockStatsServer) SetGmsAvailable(gmsAvailable int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGmsAvailable", gmsAvailable)
}

// SetGmsAvailable indicates an expected call of SetGmsAvailable.
func (mr *MockStatsServerMockRecorder) SetGmsAvailable(gmsAvailable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGmsAvailable", reflect.TypeOf((*MockStatsServer)(nil).SetGmsAvailable), gmsAvailable)
}

// SetGmsTotal mocks base method.
func (m *MockStatsServer) SetGmsTotal(gmsTotal int) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetGmsTotal", gmsTotal)
}

// SetGmsTotal indicates an expected call of SetGmsTotal.
func (mr *MockStatsServerMockRecorder) SetGmsTotal(gmsTotal interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetGmsTotal", reflect.TypeOf((*MockStatsServer)(nil).SetGmsTotal), gmsTotal)
}

// SetTickDuration mocks base method.
func (m *MockStatsServer) SetTickDuration(tickDuration time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTickDuration", tickDuration)
}

// SetTickDuration indicates an expected call of SetTickDuration.
func (mr *MockStatsServerMockRecorder) SetTickDuration(tickDuration interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTickDuration", reflect.TypeOf((*MockStatsServer)(nil).SetTickDuration), tickDuration)
}
