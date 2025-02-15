package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib" // Use pgx driver
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbName)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		panic(fmt.Sprintf("Can't connect to database: %v", err))
	}
	defer db.Close()
	// Test connection
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("Can't connect to database: %v", err))
	} else {
		fmt.Println("Successfully connected to database! User name: \n", dbUser)
	}
}
