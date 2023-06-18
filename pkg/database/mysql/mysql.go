package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// ConnectMySQL connects to the MySQL database.
func ConnectMySQL() {
	dsn := "root:password@tcp(localhost:3306)/auction_service?parseTime=true"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MySQL database")
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	return db
}