package test

import (
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/service/imp"
	"awesomeProject1/internal/tests/mock"
	"context"
	//"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepo(ctrl)

	authService := imp.NewAuthService(mockRepo)
	ctx := context.Background()
	username, password := "test_user", "password123"

	mockRepo.EXPECT().Save(ctx, &user.UserDB{Username: username, Password: password}).Return(nil)

	err := authService.Register(ctx, username, password)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepo(ctrl)
	authService := imp.NewAuthService(mockRepo)
	ctx := context.Background()
	username, password := "test_user", "password123"
	mockUser := &user.UserDB{Username: username, Password: password}

	mockRepo.EXPECT().Get(ctx, username).Return(nil, nil)                                         // Первый вызов (пользователь отсутствует)
	mockRepo.EXPECT().Save(ctx, &user.UserDB{Username: username, Password: password}).Return(nil) // Регистрация
	mockRepo.EXPECT().Get(ctx, username).Return(mockUser, nil)                                    // Второй вызов (после регистрации)

	token, err := authService.Login(ctx, username, password)
	if err != nil || token == "" {
		t.Errorf("Expected a valid token, got error: %v", err)
	}
}
