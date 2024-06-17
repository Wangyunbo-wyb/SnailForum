package router

import (
	"SnailForum/api"
	"SnailForum/middleware"
	"github.com/gin-gonic/gin"
)

func GetCommentRoutes(router *gin.RouterGroup) {
	commentGroup := router.Group("/comment")
	{
		commentGroup.GET("/list", middleware.AuthRequired(), api.CommentListHandler) // 评论列表
		commentGroup.GET("/create", middleware.AuthRequired(), api.CommentHandler)   // 评论
	}
}
