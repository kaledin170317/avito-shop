package service

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/entities/user"
	"context"
)

type UserService interface {
	Find(ctx context.Context, username string) (user.UserDB, error)
}

type AuthService interface {
	Register(ctx context.Context, username string, password string) error
	Login(ctx context.Context, username string, password string) (string, error)
	ValidateToken(ctx context.Context, token string) (string, error)
}

type ProductService interface {
	Buy(ctx context.Context, username string, name string) error
}

type InventoryService interface {
	Get(ctx context.Context, username string) ([]inventory.InventoryResponse, error)
}

type TransactionService interface {
	Get(ctx context.Context, username string) ([]transaction.TransactionRecivedResponse, []transaction.TransactionSentResponse, error)
	Save(ctx context.Context, from string, to string, amount int) error
}
