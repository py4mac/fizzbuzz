// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// MockUseCase is a mock of UseCase interface.
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase.
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance.
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockUseCase) Process(ctx context.Context) (*domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ctx)
	ret0, _ := ret[0].(*domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockUseCaseMockRecorder) Process(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockUseCase)(nil).Process), ctx)
}

// Record mocks base method.
func (m *MockUseCase) Record(ctx context.Context, e *domain.Fizzbuz) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Record", ctx, e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Record indicates an expected call of Record.
func (mr *MockUseCaseMockRecorder) Record(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockUseCase)(nil).Record), ctx, e)
}
