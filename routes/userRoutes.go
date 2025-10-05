package routes

import (
	"golang_basic_task4/controllers"

	"github.com/gin-gonic/gin"
)

// 用户路由
func UserRoutesInit(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("/login", controllers.UserController{}.Login)
	}
}
