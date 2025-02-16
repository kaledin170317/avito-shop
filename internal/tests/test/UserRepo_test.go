package test

import (
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/tests/mock"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestUserRepo_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepo(ctrl)
	ctx := context.Background()
	username := "test_user"
	mockUser := &user.UserDB{Username: username}

	mockRepo.EXPECT().Get(ctx, username).Return(mockUser, nil)

	result, err := mockRepo.Get(ctx, username)
	if err != nil || result.Username != username {
		t.Errorf("Expected user %s, got %v, error: %v", username, result, err)
	}
}

func TestUserRepo_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepo(ctrl)
	ctx := context.Background()
	usr := &user.UserDB{Username: "test_user"}

	mockRepo.EXPECT().Save(ctx, usr).Return(nil)

	err := mockRepo.Save(ctx, usr)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
