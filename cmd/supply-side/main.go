package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/HeisenbergAbhi/auction-service/internal/supply/dao"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/handler"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/service"
	"github.com/HeisenbergAbhi/auction-service/internal/supply/storage"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Create a new MySQL database connection
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/auction_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create storage, DAO, service, and handler instances
	adSpaceStorage := storage.NewAdSpaceStorage(db)
	adSpaceDAO := dao.NewAdSpaceDAO(adSpaceStorage)
	adSpaceService := service.NewAdSpaceService(adSpaceDAO)
	adSpaceHandler := handler.NewAdSpaceHandler(adSpaceService)

	// Create HTTP router and define routes
	mux := http.NewServeMux()
	mux.HandleFunc("/adspaces", adSpaceHandler.GetAllAdSpaces)
	mux.HandleFunc("/adspaces/", adSpaceHandler.GetAdSpaceByID)
	mux.HandleFunc("/adspaces/create", adSpaceHandler.CreateAdSpace)

	// Start the HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
