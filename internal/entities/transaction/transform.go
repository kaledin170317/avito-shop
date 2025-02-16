package transaction

func TransactionResponseFromDBModel(t *TransactionDB) Transaction {
	return Transaction{
		FromUser: t.FromUser.Username,
		ToUser:   t.ToUser.Username,
		Amount:   t.Amount,
	}
}
