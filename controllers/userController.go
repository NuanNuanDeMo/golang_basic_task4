package controllers

import (
	"golang_basic_task4/auth"
	"golang_basic_task4/model"
	"golang_basic_task4/module"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct{}

// 注册
func (u UserController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "密码加密失败"})
		return
	}
	user.Password = string(hashedPassword)

	if err := module.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "保存用户失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "注册成功"})
}

// 登录
func (u UserController) Login(c *gin.Context) {
	// 登录用户
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 数据库的用户
	var dbUser model.User
	if result := module.DB.Where("username = ?", user.Username).First(&dbUser); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "查询用户失败"})
		return
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	// 生成JWT token
	token, err := auth.GenerateJWT(dbUser.Username, dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"userId": dbUser.ID,
	})

}
