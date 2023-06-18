package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HeisenbergAbhi/auction-service/internal/demand/dao"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/handler"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/service"
	"github.com/HeisenbergAbhi/auction-service/internal/demand/storage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a new MySQL database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/auction_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the bidder storage
	bidderStorage := storage.NewBidderStorage(db)

	// Create the bidder DAO
	bidderDAO := dao.NewBidderDAO(bidderStorage)

	// Create the bidder service
	bidderService := service.NewBidderService(bidderDAO)

	// Create the bidder handler
	bidderHandler := handler.NewBidderHandler(bidderService)

	// Register the HTTP routes
	http.HandleFunc("/bidders", bidderHandler.GetAllBidders)
	http.HandleFunc("/bidders/", bidderHandler.GetBidderByID)
	http.HandleFunc("/bidders/create", bidderHandler.CreateBidder)

	// Get the port from the environment variable or use the default (8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the HTTP server
	log.Printf("Server running on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
