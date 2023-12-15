package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/meifanora/web-scaper-go/config"
	"log"
	"time"
)

func NewDatabase(configuration config.Config) *sql.DB {
	dbHost := configuration.Get("DB_HOST")
	dbPort := configuration.Get("DB_PORT")
	dbUser := configuration.Get("DB_USER")
	dbPassword := configuration.Get("DB_PASSWORD")
	dbName := configuration.Get("DB_NAME")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	dbConnection, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Print("\nError connecting to the database:", configuration.Get("DB_NAME"))
		panic(err)
	}

	// Test the database connection
	err = dbConnection.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
		return nil
	}

	fmt.Println("Connect to database", configuration.Get("DB_NAME"))
	dbConnection.SetMaxIdleConns(5)
	dbConnection.SetMaxOpenConns(20)
	dbConnection.SetConnMaxLifetime(60 * time.Minute)
	dbConnection.SetConnMaxIdleTime(10 * time.Minute)

	return dbConnection
}
