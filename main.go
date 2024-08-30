package main

import (
	"github.com/edmartt/bookstatic-book-service/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dbConnectObject := database.Postgres{}
	getConn := dbConnectObject.GetConnection()

	dbConnectObject.PingDB(getConn)
}
