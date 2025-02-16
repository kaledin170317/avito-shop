package imp

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/storage"
	"context"
	"fmt"
)

func NewProductService(productRepo storage.ProductRepository, userRepo storage.UserRepository, inventoryRepo storage.InventoryRepository) *ProductServiceImp {
	return &ProductServiceImp{productRepo: productRepo, userRepo: userRepo, inventoryRepo: inventoryRepo}
}

type ProductServiceImp struct {
	userRepo      storage.UserRepository
	productRepo   storage.ProductRepository
	inventoryRepo storage.InventoryRepository
}

func (p ProductServiceImp) Buy(ctx context.Context, username string, name string) error {

	user, err := p.userRepo.Get(ctx, username)

	if err != nil {
		return err
	}

	if user == nil {
		return fmt.Errorf("user not found")
	}

	product, err := p.productRepo.Get(ctx, name)

	if err != nil {
		return err
	}

	if product == nil {
		return fmt.Errorf("product not found")
	}

	if product.Price > user.Coins {
		return fmt.Errorf("not enough coins to buy %s", product.Name)
	}

	found := false
	for i := range user.Inventory {
		if user.Inventory[i].ProductID == product.ID {
			user.Inventory[i].Quantity++
			found = true
			break
		}
	}

	if !found {
		newInventoryItem := inventory.InventoryDB{
			UserID:    user.ID,
			ProductID: product.ID,
			Quantity:  1,
		}
		user.Inventory = append(user.Inventory, newInventoryItem)
	}

	user.Coins -= product.Price
	err = p.userRepo.Save(ctx, user)

	if err != nil {
		return fmt.Errorf("failed to save user and inventoryrepo: %w", err)
	}

	return nil

}
