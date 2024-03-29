// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/py4mac/fizzbuzz/internal/fizzbuzz/domain"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockRepository) Process(ctx context.Context) (*domain.Statistics, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Process", ctx)
	ret0, _ := ret[0].(*domain.Statistics)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Process indicates an expected call of Process.
func (mr *MockRepositoryMockRecorder) Process(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockRepository)(nil).Process), ctx)
}

// Record mocks base method.
func (m *MockRepository) Record(ctx context.Context, e *domain.Fizzbuz) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Record", ctx, e)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Record indicates an expected call of Record.
func (mr *MockRepositoryMockRecorder) Record(ctx, e interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Record", reflect.TypeOf((*MockRepository)(nil).Record), ctx, e)
}
