package postgres13

import (
	"awesomeProject1/internal/entities/inventory"
	"context"
	"gorm.io/gorm"
)

func NewInventoryRepo(db *gorm.DB) *InventoryRepo {
	return &InventoryRepo{db: db}
}

type InventoryRepo struct {
	db *gorm.DB
}

func (r InventoryRepo) Get(ctx context.Context, username string) ([]*inventory.InventoryDB, error) {

	var inventorydb []*inventory.InventoryDB

	result := r.db.WithContext(ctx).
		Preload("Product").
		Joins("JOIN users ON users.id = inventory.user_id").
		Where("users.username = ?", username).
		Find(&inventorydb)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(inventorydb) == 0 {
		return []*inventory.InventoryDB{}, nil
	}

	return inventorydb, nil
}

func (r InventoryRepo) Save(ctx context.Context, inventory *inventory.InventoryDB) error {
	result := r.db.WithContext(ctx).Save(inventory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
