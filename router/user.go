package router

import (
	"SnailForum/api"
	"SnailForum/middleware"
	"github.com/gin-gonic/gin"
)

func GetUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/info", middleware.AuthRequired(), api.GetUserInfoByID) //获取其他用户信息
		userGroup.POST("/register", api.Register)                              // 用户注册
		userGroup.POST("/login", api.Login)                                    // 用户登录
	}
}
