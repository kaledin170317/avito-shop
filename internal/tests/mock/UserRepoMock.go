package mocks

import (
	"awesomeProject1/internal/entities/user"
	"context"
	"github.com/golang/mock/gomock"
	"reflect"
)

type MockUserRepo struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepoMockRecorder
}

type MockUserRepoMockRecorder struct {
	mock *MockUserRepo
}

func NewMockUserRepo(ctrl *gomock.Controller) *MockUserRepo {
	mock := &MockUserRepo{ctrl: ctrl}
	mock.recorder = &MockUserRepoMockRecorder{mock}
	return mock
}

func (m *MockUserRepo) EXPECT() *MockUserRepoMockRecorder {
	return m.recorder
}

func (m *MockUserRepo) Get(ctx context.Context, username string) (*user.UserDB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, username)
	ret0, _ := ret[0].(*user.UserDB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockUserRepo) Save(ctx context.Context, user *user.UserDB) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, user)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUserRepoMockRecorder) Get(ctx, username interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUserRepo)(nil).Get), ctx, username)
}

func (mr *MockUserRepoMockRecorder) Save(ctx, user interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserRepo)(nil).Save), ctx, user)
}
