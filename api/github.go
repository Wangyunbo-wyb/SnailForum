package api

import (
	"SnailForum/common"
	"SnailForum/logic"
	"SnailForum/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GithubTrendingHandler 获取Github热榜项目
func GithubTrendingHandler(c *gin.Context) {
	p := &model.ParamGithubTrending{}
	if err := c.ShouldBindQuery(p); err != nil {
		zap.L().Error("GithubTrendingHandler with invalid params", zap.Error(err))
		panic(common.NewCustomError(common.CodeInvalidParams))
	}
	// 获取数据
	data, err := logic.GetGithubTrending(p)
	if err != nil {
		panic(common.NewCustomError(common.CodeServerBusy))
	}
	common.Success(c, data)
}
