package handlers

import (
	"awesomeProject1/internal/entities/user"
	"awesomeProject1/internal/service/imp"
	"awesomeProject1/internal/utils"
	"context"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	AuthService *imp.AuthServiceImp
}

func NewAuthHandler(authService *imp.AuthServiceImp) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Auth(w http.ResponseWriter, r *http.Request) {
	var request user.CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request")
		return
	}

	jwt, err := h.AuthService.Login(context.Background(), request.Username, request.Password)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid username or password")
		return
	}

	utils.SendSuccessResponse(w, http.StatusOK, map[string]string{"token": jwt})
}
