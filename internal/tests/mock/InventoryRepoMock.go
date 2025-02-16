package mocks

import (
	"awesomeProject1/internal/entities/inventory"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
	_ "reflect"
)

type MockInventoryRepo struct {
	ctrl     *gomock.Controller
	recorder *MockInventoryRepoMockRecorder
}

type MockInventoryRepoMockRecorder struct {
	mock *MockInventoryRepo
}

func NewMockInventoryRepo(ctrl *gomock.Controller) *MockInventoryRepo {
	mock := &MockInventoryRepo{ctrl: ctrl}
	mock.recorder = &MockInventoryRepoMockRecorder{mock}
	return mock
}

func (m *MockInventoryRepo) EXPECT() *MockInventoryRepoMockRecorder {
	return m.recorder
}

func (m *MockInventoryRepo) Get(ctx context.Context, username string) ([]*inventory.InventoryDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username)
	ret0, _ := ret[0].([]*inventory.InventoryDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockInventoryRepo) Save(ctx context.Context, inv *inventory.InventoryDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, inv)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockInventoryRepoMockRecorder) Get(ctx, username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInventoryRepo)(nil).Get), ctx, username)
}

func (mr *MockInventoryRepoMockRecorder) Save(ctx, inv interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockInventoryRepo)(nil).Save), ctx, inv)
}
