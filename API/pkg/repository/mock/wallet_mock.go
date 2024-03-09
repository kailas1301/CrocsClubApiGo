// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/repository/interfaces/wallet.go

// Package mock is a generated GoMock package.
package mock

import (
	models "CrocsClub/pkg/utils/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWalletRepository is a mock of WalletRepository interface.
type MockWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockWalletRepositoryMockRecorder
}

// MockWalletRepositoryMockRecorder is the mock recorder for MockWalletRepository.
type MockWalletRepositoryMockRecorder struct {
	mock *MockWalletRepository
}

// NewMockWalletRepository creates a new mock instance.
func NewMockWalletRepository(ctrl *gomock.Controller) *MockWalletRepository {
	mock := &MockWalletRepository{ctrl: ctrl}
	mock.recorder = &MockWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletRepository) EXPECT() *MockWalletRepositoryMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockWalletRepository) CreateWallet(userID int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletRepositoryMockRecorder) CreateWallet(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletRepository)(nil).CreateWallet), userID)
}

// GetWallet mocks base method.
func (m *MockWalletRepository) GetWallet(userID int) (models.WalletAmount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWallet", userID)
	ret0, _ := ret[0].(models.WalletAmount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWallet indicates an expected call of GetWallet.
func (mr *MockWalletRepositoryMockRecorder) GetWallet(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWallet", reflect.TypeOf((*MockWalletRepository)(nil).GetWallet), userID)
}

// GetWalletId mocks base method.
func (m *MockWalletRepository) GetWalletId(userID int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWalletId", userID)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWalletId indicates an expected call of GetWalletId.
func (mr *MockWalletRepositoryMockRecorder) GetWalletId(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWalletId", reflect.TypeOf((*MockWalletRepository)(nil).GetWalletId), userID)
}

// IsWalletExist mocks base method.
func (m *MockWalletRepository) IsWalletExist(userID int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsWalletExist", userID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWalletExist indicates an expected call of IsWalletExist.
func (mr *MockWalletRepositoryMockRecorder) IsWalletExist(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWalletExist", reflect.TypeOf((*MockWalletRepository)(nil).IsWalletExist), userID)
}

// WalletCredited mocks base method.
func (m *MockWalletRepository) WalletCredited(walletID, OrderID int, Amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletCredited", walletID, OrderID, Amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalletCredited indicates an expected call of WalletCredited.
func (mr *MockWalletRepositoryMockRecorder) WalletCredited(walletID, OrderID, Amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletCredited", reflect.TypeOf((*MockWalletRepository)(nil).WalletCredited), walletID, OrderID, Amount)
}

// WalletDebited mocks base method.
func (m *MockWalletRepository) WalletDebited(walletID, OrderID int, Amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletDebited", walletID, OrderID, Amount)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalletDebited indicates an expected call of WalletDebited.
func (mr *MockWalletRepositoryMockRecorder) WalletDebited(walletID, OrderID, Amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletDebited", reflect.TypeOf((*MockWalletRepository)(nil).WalletDebited), walletID, OrderID, Amount)
}

// WalletHistory mocks base method.
func (m *MockWalletRepository) WalletHistory(walletId int) ([]models.WalletHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalletHistory", walletId)
	ret0, _ := ret[0].([]models.WalletHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WalletHistory indicates an expected call of WalletHistory.
func (mr *MockWalletRepositoryMockRecorder) WalletHistory(walletId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalletHistory", reflect.TypeOf((*MockWalletRepository)(nil).WalletHistory), walletId)
}
