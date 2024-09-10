package main

import (
	"log"

	"github.com/edmartt/bookstatic-book-service/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("ERROR ENV LOAD: %v", err)
	}
	dbConnectObject := database.SQLite{}
	getConn := dbConnectObject.GetConnection()

	database.PingDB(getConn)
}
