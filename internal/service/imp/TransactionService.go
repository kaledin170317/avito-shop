package imp

import (
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/storage"
	"context"
	"fmt"
	"time"
)

type TransactionServiceImp struct {
	rep storage.Repository
}

func NewTransactionServiceImp(rep storage.Repository) *TransactionServiceImp {
	return &TransactionServiceImp{rep}
}

func (s TransactionServiceImp) Get(ctx context.Context, username string) ([]transaction.TransactionRecivedResponse, []transaction.TransactionSentResponse, error) {

	transactions, err := s.rep.TransactionRepository.Get(ctx, username)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get transactions: %w", err)
	}

	var received []transaction.TransactionRecivedResponse
	var sent []transaction.TransactionSentResponse

	for _, t := range transactions {
		if t.ToUser == username {
			received = append(received, transaction.TransactionRecivedResponse{
				FromUser: t.FromUser,
				Amount:   t.Amount,
			})
		}
		if t.FromUser == username {
			sent = append(sent, transaction.TransactionSentResponse{
				ToUser: t.ToUser,
				Amount: t.Amount,
			})
		}
	}

	return received, sent, nil
}

func (s TransactionServiceImp) Save(ctx context.Context, from string, to string, amount int) error {

	fromUser, err := s.rep.UserRepository.Get(ctx, from)
	if err != nil || fromUser == nil {
		return fmt.Errorf("sender not found: %s", from)
	}

	toUser, err := s.rep.UserRepository.Get(ctx, to)
	if err != nil || toUser == nil {
		return fmt.Errorf("recipient not found: %s", to)
	}

	if from == to {
		fromUser = toUser
	}

	if fromUser.Coins < amount {
		return fmt.Errorf("insufficient balance for user: %s", from)
	}

	newTransaction := &transaction.TransactionDB{
		FromUserID: fromUser.ID,
		ToUserID:   toUser.ID,
		Amount:     amount,
		CreatedAt:  time.Now(),
	}

	err = s.rep.TransactionRepository.Save(ctx, newTransaction)
	if err != nil {
		return fmt.Errorf("failed to save transaction: %w", err)
	}

	fromUser.Coins -= amount
	toUser.Coins += amount

	if err := s.rep.UserRepository.Save(ctx, fromUser); err != nil {
		return fmt.Errorf("failed to update sender balance: %w", err)
	}

	if err := s.rep.UserRepository.Save(ctx, toUser); err != nil {
		return fmt.Errorf("failed to update recipient balance: %w", err)
	}

	return nil
}
