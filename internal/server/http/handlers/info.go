package handlers

import (
	"awesomeProject1/internal/entities/info"
	inventory2 "awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/service"
	"awesomeProject1/internal/utils"
	"context"
	"net/http"
)

type InfoHandler struct {
	UserService        service.UserService
	InventoryService   service.InventoryService
	TransactionService service.TransactionService
}

func NewInfoHandler(userService service.UserService, inventoryService service.InventoryService, transactionService service.TransactionService) *InfoHandler {
	return &InfoHandler{
		UserService:        userService,
		InventoryService:   inventoryService,
		TransactionService: transactionService,
	}
}

func (h *InfoHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	user, err := h.UserService.Find(context.Background(), username)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to get user information")
		return
	}

	inventory, err := h.InventoryService.Get(context.Background(), username)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to get inventory")
		return
	}

	received, sent, err := h.TransactionService.Get(context.Background(), username)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to get transaction history")
		return
	}

	if inventory == nil {
		inventory = []inventory2.InventoryResponse{}
	}
	if received == nil {
		received = []transaction.TransactionRecivedResponse{}
	}
	if sent == nil {
		sent = []transaction.TransactionSentResponse{}
	}

	response := info.InfoResponse{
		Coins:     user.Coins,
		Inventory: inventory,
		CoinHistory: info.CoinHistory{
			Received: received,
			Sent:     sent,
		},
	}

	utils.SendSuccessResponse(w, http.StatusOK, response)
}
