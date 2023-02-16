package handler

import (
	"fmt"
	"net/http"
	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handlerBook struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *handlerBook {
	return &handlerBook{bookService}
}

func (h *handlerBook) FindAll(c *gin.Context) {
	books, err := h.bookService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (h *handlerBook) FindID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func (h *handlerBook) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	err := c.ShouldBindJSON(&bookRequest)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	book, err := h.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
