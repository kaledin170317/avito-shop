package info

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/transaction"
)

type InfoResponse struct {
	Coins       int                           `json:"coins"`
	Inventory   []inventory.InventoryResponse `json:"inventory"`
	CoinHistory CoinHistory                   `json:"coinHistory"`
}

type CoinHistory struct {
	Received []transaction.TransactionRecivedResponse `json:"received"`
	Sent     []transaction.TransactionSentResponse    `json:"sent"`
}
