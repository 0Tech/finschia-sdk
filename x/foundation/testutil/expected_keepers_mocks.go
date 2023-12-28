// Code generated by MockGen. DO NOT EDIT.
// Source: x/foundation/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/params/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthKeeper is a mock of AuthKeeper interface.
type MockAuthKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAuthKeeperMockRecorder
}

// MockAuthKeeperMockRecorder is the mock recorder for MockAuthKeeper.
type MockAuthKeeperMockRecorder struct {
	mock *MockAuthKeeper
}

// NewMockAuthKeeper creates a new mock instance.
func NewMockAuthKeeper(ctrl *gomock.Controller) *MockAuthKeeper {
	mock := &MockAuthKeeper{ctrl: ctrl}
	mock.recorder = &MockAuthKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthKeeper) EXPECT() *MockAuthKeeperMockRecorder {
	return m.recorder
}

// GetModuleAccount mocks base method.
func (m *MockAuthKeeper) GetModuleAccount(ctx context.Context, name string) types.ModuleAccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetModuleAccount", ctx, name)
	ret0, _ := ret[0].(types.ModuleAccountI)
	return ret0
}

// GetModuleAccount indicates an expected call of GetModuleAccount.
func (mr *MockAuthKeeperMockRecorder) GetModuleAccount(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetModuleAccount", reflect.TypeOf((*MockAuthKeeper)(nil).GetModuleAccount), ctx, name)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx context.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx context.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}

// MockParamsKeeper is a mock of ParamsKeeper interface.
type MockParamsKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockParamsKeeperMockRecorder
}

// MockParamsKeeperMockRecorder is the mock recorder for MockParamsKeeper.
type MockParamsKeeperMockRecorder struct {
	mock *MockParamsKeeper
}

// NewMockParamsKeeper creates a new mock instance.
func NewMockParamsKeeper(ctrl *gomock.Controller) *MockParamsKeeper {
	mock := &MockParamsKeeper{ctrl: ctrl}
	mock.recorder = &MockParamsKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamsKeeper) EXPECT() *MockParamsKeeperMockRecorder {
	return m.recorder
}

// GetSubspace mocks base method.
func (m *MockParamsKeeper) GetSubspace(s string) (types0.Subspace, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubspace", s)
	ret0, _ := ret[0].(types0.Subspace)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetSubspace indicates an expected call of GetSubspace.
func (mr *MockParamsKeeperMockRecorder) GetSubspace(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubspace", reflect.TypeOf((*MockParamsKeeper)(nil).GetSubspace), s)
}

// Subspace mocks base method.
func (m *MockParamsKeeper) Subspace(s string) types0.Subspace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subspace", s)
	ret0, _ := ret[0].(types0.Subspace)
	return ret0
}

// Subspace indicates an expected call of Subspace.
func (mr *MockParamsKeeperMockRecorder) Subspace(s interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subspace", reflect.TypeOf((*MockParamsKeeper)(nil).Subspace), s)
}
