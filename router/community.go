package router

import (
	"SnailForum/api"
	"github.com/gin-gonic/gin"
)

func GetCommunityRoutes(router *gin.RouterGroup) {
	communityGroup := router.Group("/community")
	{
		communityGroup.GET("/category", api.GetCommunityCategory) //获取社区分类
		communityGroup.GET("/:id", api.CommunityDetailHandler)    // 根据ID查找社区详情
		communityGroup.GET("/list", api.CommunityHandler)         // 获取分类社区列表
	}
}
