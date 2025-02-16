package product

import (
	"fmt"
	"gorm.io/gorm"
)

type ProductDB struct {
	ID    int    `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:name;unique;not null"`
	Price int    `gorm:"column:price;not null"`
}

func (ProductDB) TableName() string {
	return "products"
}

func InitProducts(db *gorm.DB) error {
	products := []ProductDB{
		{Name: "t-shirt", Price: 80},
		{Name: "cup", Price: 20},
		{Name: "book", Price: 50},
		{Name: "pen", Price: 10},
		{Name: "powerbank", Price: 200},
		{Name: "hoody", Price: 300},
		{Name: "umbrella", Price: 200},
		{Name: "socks", Price: 10},
		{Name: "wallet", Price: 50},
		{Name: "pink-hoody", Price: 500},
	}

	for _, product := range products {
		if err := db.FirstOrCreate(&product, ProductDB{Name: product.Name}).Error; err != nil {
			return fmt.Errorf("failed to seed product %s: %w", product.Name, err)
		}
	}

	return nil
}
