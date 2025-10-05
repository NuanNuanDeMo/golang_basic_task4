package main

import (
	"golang_basic_task4/middle"
	"golang_basic_task4/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的路由引擎
	r := gin.Default()

	// 日志
	r.Use(middle.Logger())

	// 路由
	routes.ArticleRoutesInit(r)
	routes.CommentsRoutesInit(r)
	routes.UserRoutesInit(r)

	r.Run(":8000")
}
