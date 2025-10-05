package routes

import (
	"golang_basic_task4/controllers"
	"golang_basic_task4/middle"

	"github.com/gin-gonic/gin"
)

// 评论路由
func CommentsRoutesInit(r *gin.Engine) {
	comment := r.Group("/comments", middle.AuthMiddleware())
	{
		comment.POST("/create", controllers.CommentsController{}.Create)
		comment.GET("/query/:id", controllers.CommentsController{}.Query)
	}
}
