package api

import (
	"SnailForum/common"
	"github.com/gin-gonic/gin"
)

func GetCommunityCategory(ctx *gin.Context) {
	category := logic.GetCommunityCategory()
	common.Success(ctx, category)
}
