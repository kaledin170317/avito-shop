package test

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/product"
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/service/imp"
	"awesomeProject1/internal/tests/mock"
	"context"
	//"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestProductService_Buy_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepo(ctrl)
	mockProductRepo := mocks.NewMockProductRepo(ctrl)
	mockInventoryRepo := mocks.NewMockInventoryRepo(ctrl)

	productService := imp.NewProductService(mockProductRepo, mockUserRepo, mockInventoryRepo)
	ctx := context.Background()

	username := "test_user"
	productName := "test_product"

	mockUser := &user.UserDB{ID: 1, Username: username, Coins: 100, Inventory: []inventory.InventoryDB{}}
	mockProduct := &product.ProductDB{ID: 1, Name: productName, Price: 50}

	mockUserRepo.EXPECT().Get(ctx, username).Return(mockUser, nil)
	mockProductRepo.EXPECT().Get(ctx, productName).Return(mockProduct, nil)
	mockUserRepo.EXPECT().Save(ctx, gomock.Any()).Return(nil)

	err := productService.Buy(ctx, username, productName)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestProductService_Buy_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepo(ctrl)
	mockProductRepo := mocks.NewMockProductRepo(ctrl)
	mockInventoryRepo := mocks.NewMockInventoryRepo(ctrl)

	productService := imp.NewProductService(mockProductRepo, mockUserRepo, mockInventoryRepo)
	ctx := context.Background()

	username := "test_user"
	productName := "test_product"

	mockUserRepo.EXPECT().Get(ctx, username).Return(nil, nil)

	err := productService.Buy(ctx, username, productName)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("Expected 'user not found' error, got %v", err)
	}
}

func TestProductService_Buy_ProductNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepo(ctrl)
	mockProductRepo := mocks.NewMockProductRepo(ctrl)
	mockInventoryRepo := mocks.NewMockInventoryRepo(ctrl)

	productService := imp.NewProductService(mockProductRepo, mockUserRepo, mockInventoryRepo)
	ctx := context.Background()

	username := "test_user"
	productName := "test_product"

	mockUser := &user.UserDB{ID: 1, Username: username, Coins: 100, Inventory: []inventory.InventoryDB{}}
	mockUserRepo.EXPECT().Get(ctx, username).Return(mockUser, nil)
	mockProductRepo.EXPECT().Get(ctx, productName).Return(nil, nil)

	err := productService.Buy(ctx, username, productName)
	if err == nil || err.Error() != "product not found" {
		t.Errorf("Expected 'product not found' error, got %v", err)
	}
}

func TestProductService_Buy_NotEnoughCoins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepo(ctrl)
	mockProductRepo := mocks.NewMockProductRepo(ctrl)
	mockInventoryRepo := mocks.NewMockInventoryRepo(ctrl)

	productService := imp.NewProductService(mockProductRepo, mockUserRepo, mockInventoryRepo)
	ctx := context.Background()

	username := "test_user"
	productName := "test_product"

	mockUser := &user.UserDB{ID: 1, Username: username, Coins: 30, Inventory: []inventory.InventoryDB{}}
	mockProduct := &product.ProductDB{ID: 1, Name: productName, Price: 50}

	mockUserRepo.EXPECT().Get(ctx, username).Return(mockUser, nil)
	mockProductRepo.EXPECT().Get(ctx, productName).Return(mockProduct, nil)

	err := productService.Buy(ctx, username, productName)
	if err == nil || err.Error() != "not enough coins to buy test_product" {
		t.Errorf("Expected 'not enough coins' error, got %v", err)
	}
}
