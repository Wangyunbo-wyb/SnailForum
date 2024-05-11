package api

import (
	"SnailForum/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPostList 获取帖子列表
func GetPostList(ctx *gin.Context) {
	query := dto.PostListQuery{Page: 1, PageSize: 10, Order: "create_time"}
	if err := ctx.ShouldBind(&query); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	postList := logic.GetPostList(query.Page, query.PageSize, query.Order)
	common.Success(ctx, postList)
}

// GetPostDetail 获取帖子详情
func GetPostDetail(ctx *gin.Context) {
	pidStr := ctx.Param("post_id")
	if pidStr == "" {
		zap.L().Error("post_id")
		common.FailByMsg(ctx, "post_id为空")
		return
	}
	pid, err := strconv.ParseInt(pidStr, 10, 32)
	if err != nil {

	}
	detail := logic.GetPostDetail(int32(pid))
	category := logic.GetCategoryById(detail.CategoryID)
	author := logic.GetUserById(detail.AuthorID)
	common.Success(ctx,
		vo.PostDetail{AuthorName: author.Username,
			CategoryName: category.Name,
			Post:         detail})
}

func CreatePost(ctx *gin.Context) {
	var postDTO dto.PostDTO
	if err := ctx.ShouldBind(&postDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	userId, exists := ctx.Get("userId")
	if !exists {
		fmt.Println("未登录")
	}
	logic.CreatePost(userId.(int64), postDTO)
	common.Success(ctx, nil)
}

func PostVoting(ctx *gin.Context) {
	var voteDTO dto.VoteDTO
	if err := ctx.ShouldBind(&voteDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	userId, _ := ctx.Get("userId")
	logic.PostVoting(userId.(int64), voteDTO)
	common.Success(ctx, nil)
}
