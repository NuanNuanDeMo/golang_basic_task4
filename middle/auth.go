package middle

import (
	"golang_basic_task4/auth"

	"github.com/gin-gonic/gin"
)

// 登录验证
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		token, err := auth.ParseJwt(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token"})
			return
		}
		c.Set("user_id", token.ID)
		c.Next()
	}
}
