package transaction

type TransactionSentResponse struct {
	ToUser string `json:"toUser;omitempty"`
	Amount int    `json:"amount;omitempty"`
}

type TransactionRecivedResponse struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type TransactionRequest struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}
