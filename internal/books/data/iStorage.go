package data

import "github.com/edmartt/bookstatic-book-service/internal/books/models"

type iDataAccessLayer interface {
	Create(book models.Books) (string, error)
	Read(id string) (models.Books, error)
	Update(id string) (models.Books, error)
}
