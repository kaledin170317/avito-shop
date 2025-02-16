package mocks

import (
	"awesomeProject1/internal/entities/transaction"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockTransactionRepo struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepoMockRecorder
}

type MockTransactionRepoMockRecorder struct {
	mock *MockTransactionRepo
}

func NewMockTransactionRepo(ctrl *gomock.Controller) *MockTransactionRepo {
	mock := &MockTransactionRepo{ctrl: ctrl}
	mock.recorder = &MockTransactionRepoMockRecorder{mock}
	return mock
}

func (m *MockTransactionRepo) EXPECT() *MockTransactionRepoMockRecorder {
	return m.recorder
}

func (m *MockTransactionRepo) Get(ctx context.Context, username string) ([]transaction.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username)
	ret0, _ := ret[0].([]transaction.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockTransactionRepo) Save(ctx context.Context, transaction *transaction.TransactionDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, transaction)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockTransactionRepoMockRecorder) Get(ctx, username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTransactionRepo)(nil).Get), ctx, username)
}

func (mr *MockTransactionRepoMockRecorder) Save(ctx, transaction interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockTransactionRepo)(nil).Save), ctx, transaction)
}
