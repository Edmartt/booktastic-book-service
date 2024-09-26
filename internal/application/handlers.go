package application

import (
	"net/http"

	"github.com/edmartt/bookstatic-book-service/internal/books/data"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	bookRepository data.IDataAccessLayer
}

func NewHandler(bookRepo data.IDataAccessLayer) *HTTPHandler {
	return &HTTPHandler{
		bookRepository: bookRepo,
	}
}

func (h HTTPHandler) ReadBook(context *gin.Context) {
	id := context.Param("id")
	response := httpResponse{}

	if id == "" {
		response.Response = "bad request"
		context.JSON(http.StatusBadRequest, response)
	}

	book, err := h.bookRepository.Read(id)

	if err != nil {
		response.Response = "not found"
		context.JSON(http.StatusNotFound, response)
		return
	}

	context.JSON(http.StatusOK, book)
}

func (h HTTPHandler) createBook(context *gin.Context) {

}

func (h HTTPHandler) updateBook(context *gin.Context) {

}
