package transaction

import (
	"awesomeProject1/internal/entities/user"
	"time"
)

type TransactionDB struct {
	ID         int       `gorm:"column:id;primaryKey"`
	FromUserID uint      `gorm:"column:from_user_id"`
	ToUserID   uint      `gorm:"column:to_user_id"`
	Amount     int       `gorm:"column:amount"`
	CreatedAt  time.Time `gorm:"column:created_at"`

	FromUser user.UserDB `gorm:"foreignKey:FromUserID;constraint:OnDelete:SET NULL"`
	ToUser   user.UserDB `gorm:"foreignKey:ToUserID;constraint:OnDelete:SET NULL"`
}

func (TransactionDB) TableName() string {
	return "transactions"
}

type Transaction struct {
	FromUser string `json:"fromUser"`
	ToUser   string `json:"toUser"`
	Amount   int    `json:"amount"`
}
