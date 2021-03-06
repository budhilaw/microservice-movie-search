// Code generated by MockGen. DO NOT EDIT.
// Source: ./domain/omdb.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	context "context"
	domain "microservice-test/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockOmdbUsecase is a mock of OmdbUsecase interface.
type MockOmdbUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockOmdbUsecaseMockRecorder
}

// MockOmdbUsecaseMockRecorder is the mock recorder for MockOmdbUsecase.
type MockOmdbUsecaseMockRecorder struct {
	mock *MockOmdbUsecase
}

// NewMockOmdbUsecase creates a new mock instance.
func NewMockOmdbUsecase(ctrl *gomock.Controller) *MockOmdbUsecase {
	mock := &MockOmdbUsecase{ctrl: ctrl}
	mock.recorder = &MockOmdbUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOmdbUsecase) EXPECT() *MockOmdbUsecaseMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockOmdbUsecase) Get(c context.Context, id string) (*domain.Omdb, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", c, id)
	ret0, _ := ret[0].(*domain.Omdb)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockOmdbUsecaseMockRecorder) Get(c, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockOmdbUsecase)(nil).Get), c, id)
}

// Save mocks base method.
func (m_2 *MockOmdbUsecase) Save(c context.Context, m *domain.Omdb) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Save", c, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockOmdbUsecaseMockRecorder) Save(c, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockOmdbUsecase)(nil).Save), c, m)
}

// MockOmdbRepository is a mock of OmdbRepository interface.
type MockOmdbRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOmdbRepositoryMockRecorder
}

// MockOmdbRepositoryMockRecorder is the mock recorder for MockOmdbRepository.
type MockOmdbRepositoryMockRecorder struct {
	mock *MockOmdbRepository
}

// NewMockOmdbRepository creates a new mock instance.
func NewMockOmdbRepository(ctrl *gomock.Controller) *MockOmdbRepository {
	mock := &MockOmdbRepository{ctrl: ctrl}
	mock.recorder = &MockOmdbRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOmdbRepository) EXPECT() *MockOmdbRepositoryMockRecorder {
	return m.recorder
}

// FindByImdbID mocks base method.
func (m *MockOmdbRepository) FindByImdbID(ctx context.Context, imdbId string) (*domain.Omdb, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByImdbID", ctx, imdbId)
	ret0, _ := ret[0].(*domain.Omdb)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByImdbID indicates an expected call of FindByImdbID.
func (mr *MockOmdbRepositoryMockRecorder) FindByImdbID(ctx, imdbId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByImdbID", reflect.TypeOf((*MockOmdbRepository)(nil).FindByImdbID), ctx, imdbId)
}

// Store mocks base method.
func (m_2 *MockOmdbRepository) Store(ctx context.Context, m *domain.Omdb) error {
	m_2.ctrl.T.Helper()
	ret := m_2.ctrl.Call(m_2, "Store", ctx, m)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockOmdbRepositoryMockRecorder) Store(ctx, m interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockOmdbRepository)(nil).Store), ctx, m)
}
