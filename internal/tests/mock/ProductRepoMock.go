package mocks

import (
	"awesomeProject1/internal/entities/product"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockProductRepo struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepoMockRecorder
}

type MockProductRepoMockRecorder struct {
	mock *MockProductRepo
}

func NewMockProductRepo(ctrl *gomock.Controller) *MockProductRepo {
	mock := &MockProductRepo{ctrl: ctrl}
	mock.recorder = &MockProductRepoMockRecorder{mock}
	return mock
}

func (m *MockProductRepo) EXPECT() *MockProductRepoMockRecorder {
	return m.recorder
}

func (m *MockProductRepo) Get(ctx context.Context, name string) (*product.ProductDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, name)
	ret0, _ := ret[0].(*product.ProductDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockProductRepo) Save(ctx context.Context, product product.ProductDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockProductRepoMockRecorder) Get(ctx, name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockProductRepo)(nil).Get), ctx, name)
}

func (mr *MockProductRepoMockRecorder) Save(ctx, product interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockProductRepo)(nil).Save), ctx, product)
}
