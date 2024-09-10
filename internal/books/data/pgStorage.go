package data

import (
	"fmt"

	"github.com/edmartt/bookstatic-book-service/internal/books/models"
	"github.com/edmartt/bookstatic-book-service/internal/database"
)

var dbConnectionObject database.IDBConnection

type BookDataAccess struct {
	db   database.IDBConnection
	book models.Books
}

func init() {
	dbConnectionObject = database.SQLite{}
}

func (b BookDataAccess) Create(book models.Books) (string, error) {
	conn := dbConnectionObject.GetConnection()

	dbResponse, err := conn.NamedExec("INSERT INTO books (id, isbn, title, pages, curren_page, author, year, status) VALUES(:id, :isbn, :title, :pages, :curren_page, :author, :year, :status)", &book)

	if err != nil {
		return "", err
	}

	lastID, err := dbResponse.LastInsertId()

	if err != nil {
		return "", err
	}

	bookData := fmt.Sprintf("%d %s %s", lastID, book.ISBN, book.Title)

	return bookData, nil
}

func (b BookDataAccess) Read(id string) (*models.Books, error) {
	conn := dbConnectionObject.GetConnection()

	err := conn.Get(&b.book, "SELECT isbn, title, pages, current_page, author, year, status FROM books WHERE id = ?", id)

	if err != nil {
		return nil, err
	}
	return &b.book, nil
}

func (b BookDataAccess) Update(id string, book models.Books) (string, error) {
	conn := dbConnectionObject.GetConnection()

	result, err := conn.NamedExec("UPDATE books SET isbn = :isbn, title = :title, pages = :pages, current_page = :current_page, author = :author, year = :year, status = :status WHERE id=?", book)

	if err != nil {
		return "", err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return "", err
	}

	bookData := fmt.Sprintf("%d %s %s", lastID, book.ISBN, book.Title)

	return bookData, nil
}
