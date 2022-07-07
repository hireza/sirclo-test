// Code generated by MockGen. DO NOT EDIT.
// Source: domain/weight.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/hireza/sirclo-test/berat/domain"
)

// MockWeightRepository is a mock of WeightRepository interface.
type MockWeightRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWeightRepositoryMockRecorder
}

// MockWeightRepositoryMockRecorder is the mock recorder for MockWeightRepository.
type MockWeightRepositoryMockRecorder struct {
	mock *MockWeightRepository
}

// NewMockWeightRepository creates a new mock instance.
func NewMockWeightRepository(ctrl *gomock.Controller) *MockWeightRepository {
	mock := &MockWeightRepository{ctrl: ctrl}
	mock.recorder = &MockWeightRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeightRepository) EXPECT() *MockWeightRepositoryMockRecorder {
	return m.recorder
}

// CheckByDate mocks base method.
func (m *MockWeightRepository) CheckByDate(ctx context.Context, date string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckByDate", ctx, date)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckByDate indicates an expected call of CheckByDate.
func (mr *MockWeightRepositoryMockRecorder) CheckByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckByDate", reflect.TypeOf((*MockWeightRepository)(nil).CheckByDate), ctx, date)
}

// Create mocks base method.
func (m *MockWeightRepository) Create(ctx context.Context, data *domain.Weight) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWeightRepositoryMockRecorder) Create(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWeightRepository)(nil).Create), ctx, data)
}

// Delete mocks base method.
func (m *MockWeightRepository) Delete(ctx context.Context, date string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, date)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWeightRepositoryMockRecorder) Delete(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWeightRepository)(nil).Delete), ctx, date)
}

// GetAll mocks base method.
func (m *MockWeightRepository) GetAll(ctx context.Context) ([]*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWeightRepositoryMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWeightRepository)(nil).GetAll), ctx)
}

// GetByDate mocks base method.
func (m *MockWeightRepository) GetByDate(ctx context.Context, date string) (*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", ctx, date)
	ret0, _ := ret[0].(*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockWeightRepositoryMockRecorder) GetByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockWeightRepository)(nil).GetByDate), ctx, date)
}

// Update mocks base method.
func (m *MockWeightRepository) Update(ctx context.Context, date string, data *domain.Weight) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, date, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWeightRepositoryMockRecorder) Update(ctx, date, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWeightRepository)(nil).Update), ctx, date, data)
}

// MockWeightUsecase is a mock of WeightUsecase interface.
type MockWeightUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockWeightUsecaseMockRecorder
}

// MockWeightUsecaseMockRecorder is the mock recorder for MockWeightUsecase.
type MockWeightUsecaseMockRecorder struct {
	mock *MockWeightUsecase
}

// NewMockWeightUsecase creates a new mock instance.
func NewMockWeightUsecase(ctrl *gomock.Controller) *MockWeightUsecase {
	mock := &MockWeightUsecase{ctrl: ctrl}
	mock.recorder = &MockWeightUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeightUsecase) EXPECT() *MockWeightUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWeightUsecase) Create(ctx context.Context, data *domain.Weight) ([]*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, data)
	ret0, _ := ret[0].([]*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockWeightUsecaseMockRecorder) Create(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWeightUsecase)(nil).Create), ctx, data)
}

// Delete mocks base method.
func (m *MockWeightUsecase) Delete(ctx context.Context, date string) ([]*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, date)
	ret0, _ := ret[0].([]*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockWeightUsecaseMockRecorder) Delete(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWeightUsecase)(nil).Delete), ctx, date)
}

// GetAll mocks base method.
func (m *MockWeightUsecase) GetAll(ctx context.Context) ([]*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx)
	ret0, _ := ret[0].([]*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWeightUsecaseMockRecorder) GetAll(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWeightUsecase)(nil).GetAll), ctx)
}

// GetByDate mocks base method.
func (m *MockWeightUsecase) GetByDate(ctx context.Context, date string) (*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate", ctx, date)
	ret0, _ := ret[0].(*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockWeightUsecaseMockRecorder) GetByDate(ctx, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockWeightUsecase)(nil).GetByDate), ctx, date)
}

// Update mocks base method.
func (m *MockWeightUsecase) Update(ctx context.Context, date string, data *domain.Weight) ([]*domain.Weight, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, date, data)
	ret0, _ := ret[0].([]*domain.Weight)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockWeightUsecaseMockRecorder) Update(ctx, date, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWeightUsecase)(nil).Update), ctx, date, data)
}

// MockWeightHandler is a mock of WeightHandler interface.
type MockWeightHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWeightHandlerMockRecorder
}

// MockWeightHandlerMockRecorder is the mock recorder for MockWeightHandler.
type MockWeightHandlerMockRecorder struct {
	mock *MockWeightHandler
}

// NewMockWeightHandler creates a new mock instance.
func NewMockWeightHandler(ctrl *gomock.Controller) *MockWeightHandler {
	mock := &MockWeightHandler{ctrl: ctrl}
	mock.recorder = &MockWeightHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeightHandler) EXPECT() *MockWeightHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockWeightHandler) Create() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockWeightHandlerMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockWeightHandler)(nil).Create))
}

// Delete mocks base method.
func (m *MockWeightHandler) Delete() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockWeightHandlerMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockWeightHandler)(nil).Delete))
}

// GetAll mocks base method.
func (m *MockWeightHandler) GetAll() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockWeightHandlerMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockWeightHandler)(nil).GetAll))
}

// GetByDate mocks base method.
func (m *MockWeightHandler) GetByDate() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByDate")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// GetByDate indicates an expected call of GetByDate.
func (mr *MockWeightHandlerMockRecorder) GetByDate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByDate", reflect.TypeOf((*MockWeightHandler)(nil).GetByDate))
}

// Update mocks base method.
func (m *MockWeightHandler) Update() http.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update")
	ret0, _ := ret[0].(http.Handler)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockWeightHandlerMockRecorder) Update() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockWeightHandler)(nil).Update))
}
