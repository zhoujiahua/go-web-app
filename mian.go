package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoujiahua/go-web-app/handlers"
	"github.com/zhoujiahua/go-web-app/services"
)

func main() {
	router := gin.Default()

	bookService := services.NewBookService()
	bookHandler := handlers.NewBookHandler(bookService)

	router.GET("/books", bookHandler.GetAllBooks)
	router.GET("/books/:id", bookHandler.GetBookByID)
	router.POST("/books", bookHandler.AddBook)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)

	router.Run(":8080")
}
