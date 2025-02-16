package postgres13

import (
	"awesomeProject1/internal/entities/transaction"
	"context"
	//"fmt"
	"gorm.io/gorm"
)

func NewTransactionRepo(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

type TransactionRepository struct {
	db *gorm.DB
}

func (r TransactionRepository) Get(ctx context.Context, username string) ([]transaction.Transaction, error) {
	var transactionsDB []*transaction.TransactionDB

	result := r.db.WithContext(ctx).
		Preload("FromUser").
		Preload("ToUser").
		Joins("JOIN users u1 ON u1.id = transactions.from_user_id").
		Joins("JOIN users u2 ON u2.id = transactions.to_user_id").
		Where("u1.username = ? OR u2.username = ?", username, username).
		Find(&transactionsDB)

	if result.Error != nil {
		return nil, result.Error
	}

	if len(transactionsDB) == 0 {
		return []transaction.Transaction{}, nil
	}

	var ts []transaction.Transaction
	for _, t := range transactionsDB {
		tr := transaction.TransactionResponseFromDBModel(t)
		ts = append(ts, tr)
	}

	return ts, nil
}

func (r TransactionRepository) Save(ctx context.Context, transaction *transaction.TransactionDB) error {
	result := r.db.WithContext(ctx).Save(transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
