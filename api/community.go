package api

import (
	"SnailForum/common"
	"SnailForum/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 社区

// GetCommunityCategory 查找社区种类
func GetCommunityCategory(ctx *gin.Context) {
	category := logic.GetCommunityCategory()
	common.Success(ctx, category)
}

// CommunityHandler 查找社区列表
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区(community_id,community_name)以列表的形式返回
	communityList, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		panic(common.NewCustomError(common.CodeServerBusy)) // 不轻易把服务端报错暴露给外面
	}
	common.Success(c, communityList)
}

// CommunityDetailHandler 根据ID查找到社区分类的详情
func CommunityDetailHandler(c *gin.Context) {
	// 1、获取社区ID
	communityIdStr := c.Param("id")                               // 获取URL参数
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64) // id字符串格式转换
	if err != nil {
		panic(common.NewCustomError(common.CodeInvalidParams))
	}

	// 2、根据ID获取社区详情
	communityList, err := logic.GetCommunityByID(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID() failed", zap.Error(err))
		panic(common.NewCustomError(common.CodeFoundFailed))
	}
	common.Success(c, communityList)
}
