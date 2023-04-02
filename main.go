package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhoujiahua/go-web-app/controllers"
	"github.com/zhoujiahua/go-web-app/routes"
)

func main() {
	router := gin.Default()
	bookController := controllers.NewBookController(router)
	bookRouter := routes.NewBookRouter(router, bookController)
	bookRouter.RegisterRoutes()
	router.Run(":9527")
}
