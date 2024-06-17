package router

import (
	"SnailForum/api"
	"github.com/gin-gonic/gin"
)

func GetGithubRoutes(router *gin.RouterGroup) {
	githubGroup := router.Group("/github")
	{
		githubGroup.GET("/trending", api.GithubTrendingHandler) // Github热榜
	}
}
