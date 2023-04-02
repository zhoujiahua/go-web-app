package handlers

import (
	"net/http"
	"strconv"

	"restful/jerry/go-web-app/models"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookService models.BookService
}

func NewBookHandler(bookService models.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetAllBooks(c *gin.Context) {
	books := h.bookService.GetAllBooks()
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book = h.bookService.AddBook(book)
	c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book, err := h.bookService.UpdateBook(id, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.bookService.DeleteBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
