package router

import (
	"SnailForum/api"
	"github.com/gin-gonic/gin"
)

func GetCommunityRoutes(router *gin.RouterGroup) {
	communityGroup := router.Group("/community")
	{
		communityGroup.GET("/category", api.GetCommunityCategory) //获取社区分类
	}
}
