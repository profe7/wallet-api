package main

import (
	"log"
	"net/http"

	"wallet-api/db"
	"wallet-api/handler"
	"wallet-api/middleware"
	"wallet-api/repository"
	"wallet-api/service"
)

func main() {
	// 1. Initialise the database (like @Bean DataSource in Spring).
	database := db.InitDB("./wallet.db")
	defer database.Close()

	// 2. Create layers - manual dependency injection.
	accountRepo := repository.NewAccountRepository(database)
	walletService := service.NewWalletService(accountRepo)
	walletHandler := handler.NewWalletHandler(walletService)

	// 3. Register routes (like @RequestMapping).
	http.HandleFunc("/balance", middleware.ErrorHandler(walletHandler.CheckBalance))
	http.HandleFunc("/withdraw", middleware.ErrorHandler(walletHandler.Withdraw))

	// 4. Start the server.
	log.Println("[Server] started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
