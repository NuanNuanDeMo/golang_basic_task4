package controllers

import (
	"fmt"
	"golang_basic_task4/auth"
	"golang_basic_task4/model"
	"golang_basic_task4/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
}

// 文章新增
func (a ArticleController) Add(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := module.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "创建文章失败"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "文章创建成功"})
}

// 文章查询单个
func (a ArticleController) Query(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("dayinddd", id)
	var post model.Post
	module.DB.Preload("Comments").First(&post, id)
	c.JSON(http.StatusOK, gin.H{"post": post})
}

// 文章查询列表
func (a ArticleController) QueryList(c *gin.Context) {
	var posts []model.Post
	if err := module.DB.Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章列表失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// 文章修改
func (a ArticleController) Edit(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var existingPost model.Post
	if err := module.DB.Where("id = ?", post.ID).First(&existingPost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	user, _ := auth.ParseJwt(tokenString)

	if existingPost.UserID != user.Uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限修改"})
		return
	}
	if err := module.DB.Model(&existingPost).Updates(post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "修改失败"})
		return
	}
	c.JSON(http.StatusOK, post)
}

// 文章删除
func (a ArticleController) Delete(c *gin.Context) {
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var deletePost model.Post
	if err := module.DB.Where("id = ?", post.ID).First(&deletePost).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "没有查询到文章信息"})
		return
	}

	tokenString := c.GetHeader("Authorization")
	user, _ := auth.ParseJwt(tokenString)

	if deletePost.UserID != user.Uid {
		c.JSON(http.StatusForbidden, gin.H{"error": "你没有权限修改"})
		return
	}

	if err := module.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "文章删除失败"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "文章删除成功"})

}
