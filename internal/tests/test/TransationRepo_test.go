package test

import (
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/tests/mock"
	"context"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

func TestTransactionRepository_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepo(ctrl)
	ctx := context.Background()
	username := "test_user"
	mockTransactions := []transaction.Transaction{
		{FromUser: "Alice", ToUser: "Bob", Amount: 100},
		{FromUser: "Charlie", ToUser: "Dave", Amount: 200},
	}

	mockRepo.EXPECT().Get(ctx, username).Return(mockTransactions, nil)

	result, err := mockRepo.Get(ctx, username)
	if err != nil || len(result) != len(mockTransactions) {
		t.Errorf("Expected %d transactions, got %d, error: %v", len(mockTransactions), len(result), err)
	}
}

func TestTransactionRepository_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTransactionRepo(ctrl)
	ctx := context.Background()
	trans := &transaction.TransactionDB{
		ID:         1,
		FromUserID: 1,
		ToUserID:   2,
		Amount:     500,
		CreatedAt:  time.Now(),
	}

	mockRepo.EXPECT().Save(ctx, trans).Return(nil)

	err := mockRepo.Save(ctx, trans)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
