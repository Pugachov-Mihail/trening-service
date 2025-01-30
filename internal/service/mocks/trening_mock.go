// Code generated by MockGen. DO NOT EDIT.
// Source: trening.go
//
// Generated by this command:
//
//	mockgen -source=trening.go -destination=mocks/trening_mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	slog "log/slog"
	reflect "reflect"
	trening_v1 "trening/protos/gen/trening.v1"

	gomock "go.uber.org/mock/gomock"
)

// MockTreningStorage is a mock of TreningStorage interface.
type MockTreningStorage struct {
	ctrl     *gomock.Controller
	recorder *MockTreningStorageMockRecorder
	isgomock struct{}
}

// MockTreningStorageMockRecorder is the mock recorder for MockTreningStorage.
type MockTreningStorageMockRecorder struct {
	mock *MockTreningStorage
}

// NewMockTreningStorage creates a new mock instance.
func NewMockTreningStorage(ctrl *gomock.Controller) *MockTreningStorage {
	mock := &MockTreningStorage{ctrl: ctrl}
	mock.recorder = &MockTreningStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTreningStorage) EXPECT() *MockTreningStorageMockRecorder {
	return m.recorder
}

// AddTreningSourse mocks base method.
func (m *MockTreningStorage) AddTreningSourse(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTreningSourse", ctx, idTrening, userId, log)
	ret0, _ := ret[0].([]trening_v1.GetTreningList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTreningSourse indicates an expected call of AddTreningSourse.
func (mr *MockTreningStorageMockRecorder) AddTreningSourse(ctx, idTrening, userId, log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTreningSourse", reflect.TypeOf((*MockTreningStorage)(nil).AddTreningSourse), ctx, idTrening, userId, log)
}

// DeletedTreningUserService mocks base method.
func (m *MockTreningStorage) DeletedTreningUserService(ctx context.Context, idTrening, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletedTreningUserService", ctx, idTrening, userId, log)
	ret0, _ := ret[0].([]trening_v1.GetTreningList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeletedTreningUserService indicates an expected call of DeletedTreningUserService.
func (mr *MockTreningStorageMockRecorder) DeletedTreningUserService(ctx, idTrening, userId, log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletedTreningUserService", reflect.TypeOf((*MockTreningStorage)(nil).DeletedTreningUserService), ctx, idTrening, userId, log)
}

// GetCurrentTreningService mocks base method.
func (m *MockTreningStorage) GetCurrentTreningService(ctx context.Context, idTrening, userId int64, log *slog.Logger) (trening_v1.GetTreningList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCurrentTreningService", ctx, idTrening, userId, log)
	ret0, _ := ret[0].(trening_v1.GetTreningList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCurrentTreningService indicates an expected call of GetCurrentTreningService.
func (mr *MockTreningStorageMockRecorder) GetCurrentTreningService(ctx, idTrening, userId, log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCurrentTreningService", reflect.TypeOf((*MockTreningStorage)(nil).GetCurrentTreningService), ctx, idTrening, userId, log)
}

// GetUserTreningService mocks base method.
func (m *MockTreningStorage) GetUserTreningService(ctx context.Context, userId int64, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserTreningService", ctx, userId, log)
	ret0, _ := ret[0].([]trening_v1.GetTreningList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserTreningService indicates an expected call of GetUserTreningService.
func (mr *MockTreningStorageMockRecorder) GetUserTreningService(ctx, userId, log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserTreningService", reflect.TypeOf((*MockTreningStorage)(nil).GetUserTreningService), ctx, userId, log)
}

// TreningListSourse mocks base method.
func (m *MockTreningStorage) TreningListSourse(ctx context.Context, page, offset int32, log *slog.Logger) ([]trening_v1.GetTreningList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TreningListSourse", ctx, page, offset, log)
	ret0, _ := ret[0].([]trening_v1.GetTreningList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TreningListSourse indicates an expected call of TreningListSourse.
func (mr *MockTreningStorageMockRecorder) TreningListSourse(ctx, page, offset, log any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TreningListSourse", reflect.TypeOf((*MockTreningStorage)(nil).TreningListSourse), ctx, page, offset, log)
}
