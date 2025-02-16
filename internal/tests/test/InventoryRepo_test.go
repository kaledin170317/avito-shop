package test

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/product"
	"awesomeProject1/internal/tests/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInventoryRepo_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockInventoryRepo(ctrl)
	ctx := context.Background()
	username := "testuser"

	t.Run("Success - returns inventory", func(t *testing.T) {
		expectedInventory := []*inventory.InventoryDB{
			{ID: 1, UserID: 1, ProductID: 1, Quantity: 5, Product: product.ProductDB{Name: "T-Shirt"}},
		}
		mockRepo.EXPECT().Get(ctx, username).Return(expectedInventory, nil)

		result, err := mockRepo.Get(ctx, username)
		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "T-Shirt", result[0].Product.Name)
	})

	t.Run("Empty Inventory - returns empty list", func(t *testing.T) {
		mockRepo.EXPECT().Get(ctx, username).Return([]*inventory.InventoryDB{}, nil)

		result, err := mockRepo.Get(ctx, username)
		assert.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("DB Error - returns error", func(t *testing.T) {
		mockRepo.EXPECT().Get(ctx, username).Return(nil, errors.New("database error"))

		result, err := mockRepo.Get(ctx, username)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")
	})
}
