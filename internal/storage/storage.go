package storage

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/product"
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/storage/postgres13"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	Get(ctx context.Context, username string) (*user.UserDB, error)
	Save(ctx context.Context, user *user.UserDB) error
}

type ProductRepository interface {
	Get(ctx context.Context, name string) (*product.ProductDB, error)
	Save(ctx context.Context, product product.ProductDB) error
}

type InventoryRepository interface {
	Get(ctx context.Context, username string) ([]*inventory.InventoryDB, error)
	Save(ctx context.Context, inventory *inventory.InventoryDB) error
}

type TransactionRepository interface {
	Get(ctx context.Context, username string) ([]transaction.Transaction, error)
	Save(ctx context.Context, transaction *transaction.TransactionDB) error
}

type Repository struct {
	UserRepository
	ProductRepository
	InventoryRepository
	TransactionRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:        postgres13.NewUserRepo(db),
		ProductRepository:     postgres13.NewProductRepo(db),
		InventoryRepository:   postgres13.NewInventoryRepo(db),
		TransactionRepository: postgres13.NewTransactionRepo(db),
	}
}
