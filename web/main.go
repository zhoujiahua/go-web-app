package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 初始化控制器和路由
	bookController := controllers.NewBookController()
	bookRouter := routes.NewBookRouter(router, bookController)

	// 注册路由
	bookRouter.RegisterRoutes()

	// 运行服务器
	router.Run(":8080")
}
