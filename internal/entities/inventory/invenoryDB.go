package inventory

import (
	"awesomeProject1/internal/entities/product"
)

type InventoryDB struct {
	ID        int  `gorm:"column:id;primaryKey"`
	UserID    uint `gorm:"column:user_id;not null"`           // Foreign key to UserDB
	ProductID int  `gorm:"column:product_id;not null;unique"` // Foreign key to ProductDB
	Quantity  int  `gorm:"column:quantity;default:1"`

	Product product.ProductDB `gorm:"foreignKey:ProductID;constraint:OnDelete:RESTRICT"`
}

func (InventoryDB) TableName() string {
	return "inventory"
}
