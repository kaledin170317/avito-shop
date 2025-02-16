package postgres13

import (
	"awesomeProject1/internal/entities/product"
	"context"
	"errors"
	"gorm.io/gorm"
)

func NewProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{db: db}
}

type ProductRepo struct {
	db *gorm.DB
}

func (r ProductRepo) Get(ctx context.Context, name string) (*product.ProductDB, error) {
	var p product.ProductDB
	result := r.db.WithContext(ctx).Where("name = ?", name).First(&p)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &p, nil
}

func (r ProductRepo) Save(ctx context.Context, product product.ProductDB) error {
	result := r.db.WithContext(ctx).Save(product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
