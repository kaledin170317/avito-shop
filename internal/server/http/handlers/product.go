package handlers

import (
	"awesomeProject1/internal/service/imp"
	"awesomeProject1/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type ProductHandler struct {
	ProductService *imp.ProductServiceImp
}

func NewProductHandler(productService *imp.ProductServiceImp) *ProductHandler {
	return &ProductHandler{ProductService: productService}
}

func (h *ProductHandler) Buy(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)

	vars := mux.Vars(r)
	item := vars["item"]
	if item == "" {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Item not specified")
		return
	}

	err := h.ProductService.Buy(r.Context(), username, item)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}
