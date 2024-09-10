package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // sqlx uses it indirectly
)

type Postgres struct {
	user     string
	db       string
	password string
	host     string
	port     string
}

// GetConnection connects to specific database
func (p Postgres) GetConnection() *sqlx.DB {
	p.user = os.Getenv("PG_USER")
	p.db = os.Getenv("PG_DB")
	p.password = os.Getenv("PG_PASSWORD")
	p.host = os.Getenv("HOST")
	p.port = os.Getenv("PORT")

	db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", p.user, p.db, p.password, p.host, p.port))

	if err != nil {
		log.Fatal("CONNECTION ERROR: ", err)
	}

	return db
}
