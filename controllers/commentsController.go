package controllers

import (
	"golang_basic_task4/auth"
	"golang_basic_task4/model"
	"golang_basic_task4/module"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsController struct{}

// 新增评论
func (c CommentsController) Create(cxt *gin.Context) {
	var comment model.Comment
	if err := cxt.ShouldBindJSON(&comment); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 获取用户ID
	tokenString := cxt.GetHeader("Authorization")
	user, _ := auth.ParseJwt(tokenString)

	comment.UserID = user.Uid

	if err := module.DB.Create(&comment).Error; err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "创建评论失败"})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{"message": "评论创建成功"})
}

// 查询评论列表
func (c CommentsController) Query(cxt *gin.Context) {
	postId := cxt.Param("id")
	var comments []model.Comment
	if err := module.DB.Where("post_id = ?", postId).Find(&comments).Error; err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": "查询评论失败"})
		return
	}
	cxt.JSON(http.StatusOK, comments)
}
