package routes

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

type BookRouter interface {
	RegisterRoutes()
}

type bookRouter struct {
	bookController controllers.BookController
}

func NewBookRouter(router *gin.Engine, bookController controllers.BookController) BookRouter {
	return &bookRouter{
		bookController: bookController,
	}
}

func (r *bookRouter) RegisterRoutes() {
	bookRoutes := r.bookController.Router()

	v1 := bookRoutes.Group("/api/v1/books")
	{
		v1.GET("", r.bookController.ListBooks)
		v1.GET("/:id", r.bookController.GetBook)
		v1.POST("", r.bookController.CreateBook)
		v1.PUT("/:id", r.bookController.UpdateBook)
		v1.DELETE("/:id", r.bookController.DeleteBook)
	}
}
