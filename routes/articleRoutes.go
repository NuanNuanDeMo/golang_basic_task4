package routes

import (
	"golang_basic_task4/controllers"
	"golang_basic_task4/middle"

	"github.com/gin-gonic/gin"
)

// 文章路由
func ArticleRoutesInit(r *gin.Engine) {
	a := r.Group("/article", middle.AuthMiddleware())
	{
		a.POST("/add", controllers.ArticleController{}.Add)
		a.GET("/query/:id", controllers.ArticleController{}.Query)
		a.GET("/queryList", controllers.ArticleController{}.QueryList)
		a.POST("/edit", controllers.ArticleController{}.Edit)
		a.POST("/delete", controllers.ArticleController{}.Delete)
	}
}
