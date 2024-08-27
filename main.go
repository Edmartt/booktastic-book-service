package main

import (
	"github.com/edmartt/bookstatic-book-service/internal/database"
)

func main() {
	dbConnectObject := database.Postgres{}
	getConn := dbConnectObject.GetConnection()

	dbConnectObject.PingDB(getConn)
}
