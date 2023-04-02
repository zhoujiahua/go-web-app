package main

import (
	"restful/jerry/go-web-app/controllers"
	"restful/jerry/go-web-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	bookController := controllers.NewBookController(router)
	bookRouter := routes.NewBookRouter(router, bookController)
	bookRouter.RegisterRoutes()
	router.Run(":9527")
}
