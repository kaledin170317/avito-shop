package test

import (
	"awesomeProject1/internal/entities/product"
	mocks "awesomeProject1/internal/tests/mock"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProductRepo_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	ctx := context.Background()
	name := "test_product"

	mockRepo.EXPECT().Get(ctx, name).Return(&product.ProductDB{Name: name}, nil)

	result, err := mockRepo.Get(ctx, name)
	if err != nil || result.Name != name {
		t.Errorf("Expected product with name %s, got %v, error: %v", name, result, err)
	}
}

func TestProductRepo_Save(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockProductRepo(ctrl)
	ctx := context.Background()
	prod := product.ProductDB{Name: "new_product"}

	mockRepo.EXPECT().Save(ctx, prod).Return(nil)

	err := mockRepo.Save(ctx, prod)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
