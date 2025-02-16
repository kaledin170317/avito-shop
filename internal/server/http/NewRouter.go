package server

import (
	"awesomeProject1/internal/server/http/handlers"
	"awesomeProject1/internal/service/imp"
	"awesomeProject1/internal/storage"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"net/http"
)

func NewRouter(db *gorm.DB) *mux.Router {

	r := mux.NewRouter()

	rep := storage.NewRepository(db)

	authService := imp.NewAuthService(rep.UserRepository)
	productService := imp.NewProductService(rep.ProductRepository, rep.UserRepository, rep.InventoryRepository)
	transactionService := imp.NewTransactionServiceImp(*rep)
	inventoryService := imp.NewInventoryServiceImp(*rep)
	userService := imp.NewUserService(rep.UserRepository)

	authHandler := handlers.NewAuthHandler(authService)
	productHandler := handlers.NewProductHandler(productService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	infoHandler := handlers.NewInfoHandler(userService, inventoryService, transactionService)

	r.Use(imp.JWTMiddleware(authService))

	r.HandleFunc("/api/auth", authHandler.Auth).Methods(http.MethodPost) // Авторизация

	r.HandleFunc("/api/buy/{item}", productHandler.Buy).Methods(http.MethodGet) // Покупка товара

	r.HandleFunc("/api/sendCoin", transactionHandler.SendCoins).Methods(http.MethodPost)

	r.HandleFunc("/api/info", infoHandler.GetInfo).Methods(http.MethodGet)

	return r
}
