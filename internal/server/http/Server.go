package server

import (
	"awesomeProject1/internal/entities/inventory"
	"awesomeProject1/internal/entities/product"
	"awesomeProject1/internal/entities/transaction"
	"awesomeProject1/internal/entities/user"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"time"
)

func StartServer() {

	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	userDB := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", userDB, password, host, port, dbName)
	fmt.Println("Connecting to:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	migrate(db, &user.UserDB{})
	migrate(db, &transaction.TransactionDB{})
	migrate(db, &inventory.InventoryDB{})
	migrate(db, &product.ProductDB{})

	if err := product.InitProducts(db); err != nil {
		log.Fatalf("Failed to seed products: %v", err)
	}

	r := NewRouter(db)

	server := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(server.ListenAndServe())
}

func migrate[T any](db *gorm.DB, model *T) {
	err := db.AutoMigrate(model)
	if err != nil {
		log.Fatalf("Failed to migrate %T: %v", model, err)
	}
}
