package handlers

import (
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/service"
	"awesomeProject1/internal/utils"
	"context"
	"encoding/json"
	"net/http"
)

type TransactionHandler struct {
	TransactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

func (h *TransactionHandler) SendCoins(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	var req transaction.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err := h.TransactionService.Save(context.Background(), username, req.ToUser, req.Amount)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to send coins: "+err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
