package user

import (
	"awesomeProject1/internal/entities/inventory"
)

type UserDB struct {
	ID       uint   `gorm:"column:id;primaryKey"`
	Username string `gorm:"column:username;unique;not null"`
	Password string `gorm:"column:password;not null"`
	Coins    int    `gorm:"column:coins;default:1000"`

	Inventory []inventory.InventoryDB `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (UserDB) TableName() string {
	return "users"
}
