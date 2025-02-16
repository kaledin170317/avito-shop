package postgres13

import (
	"awesomeProject1/internal/entities/user"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r UserRepo) Get(ctx context.Context, username string) (*user.UserDB, error) {
	var foundUser user.UserDB
	result := r.db.WithContext(ctx).
		Preload("Inventory").
		Where("username = ?", username).
		First(&foundUser)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &foundUser, nil
}

func (r UserRepo) Save(ctx context.Context, user *user.UserDB) error {

	tx := r.db.WithContext(ctx).Begin()

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	for _, inventory := range user.Inventory {
		if err := tx.Save(&inventory).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
