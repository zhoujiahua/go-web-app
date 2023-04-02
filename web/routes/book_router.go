package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoujiahua/test-app/controllers"
)

type BookRouter struct {
	router         *gin.Engine
	bookController *controllers.BookController
}

func NewBookRouter(router *gin.Engine, bookController *controllers.BookController) *BookRouter {
	return &BookRouter{
		router:         router,
		bookController: bookController,
	}
}

func (r *BookRouter) RegisterRoutes() {
	booksGroup := r.router.Group("/books")
	{
		booksGroup.GET("", r.bookController.GetAllBooks)
		booksGroup.GET("/:id", r.bookController.GetBookById)
		booksGroup.POST("", r.bookController.CreateBook)
		booksGroup.PUT("/:id", r.bookController.UpdateBookById)
		booksGroup.DELETE("/:id", r.bookController.DeleteBookById)
	}
}
