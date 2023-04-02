package routes

import (
	"restful/jerry/go-web-app/controllers"

	"github.com/gin-gonic/gin"
)

type BookRouter interface {
	RegisterRoutes()
}

type bookRouter struct {
	bookController *controllers.BookController
}

func NewBookRouter(router *gin.Engine, bookController *controllers.BookController) BookRouter {
	return &bookRouter{
		bookController: bookController,
	}
}

func (r *bookRouter) RegisterRoutes() {
	bookRoutes := r.bookController.Router

	v1 := bookRoutes.Group("")
	{
		v1.GET("", r.bookController.GetBooks)
		v1.GET("/:id", r.bookController.GetBookByID)
		v1.POST("", r.bookController.AddBook)
		v1.PUT("/:id", r.bookController.UpdateBook)
		v1.DELETE("/:id", r.bookController.DeleteBook)
	}
}
