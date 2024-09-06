package main

import (
	"github.com/edmartt/bookstatic-book-service/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbConnectObject := database.SQLite{}
	getConn := dbConnectObject.GetConnection()

	database.PingDB(getConn)
}
