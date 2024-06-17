package api

import (
	"SnailForum/common"
	"SnailForum/logic"
	"SnailForum/model"
	"SnailForum/pkg/snowflake"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

//评论

// CommentHandler 创建评论
func CommentHandler(c *gin.Context) {
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		zap.L().Error("c.BindJSON() failed", zap.Error(err))
		panic(common.NewCustomError(common.CodeInvalidParams))
	}
	// 生成评论ID
	commentID := snowflake.GenerateID()
	// 获取作者ID，当前请求的UserID
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		panic(common.NewCustomError(common.CodeNotLogin))
	}
	comment.CommentID = commentID
	comment.AuthorID = userID
	// 创建评论
	if err := logic.CreateComment(&comment); err != nil {
		zap.L().Error("mysql.CreateComment(&comment) failed", zap.Error(err))
		panic(common.NewCustomError(common.CodeServerBusy))
	}
	common.Success(c, nil)
}

// CommentListHandler 评论列表
func CommentListHandler(c *gin.Context) {
	idsStr, ok := c.GetQueryArray("ids")
	if !ok {
		panic(common.NewCustomError(common.CodeInvalidParams))
	}

	// 将字符串类型的 ID 列表转换为 uint64 类型的切片
	var ids []uint64
	for _, idStr := range idsStr {
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			panic(common.NewCustomError(common.CodeInvalidParams))
		}
		ids = append(ids, id)
	}

	posts, err := logic.GetCommentListByIDs(ids)
	if err != nil {
		panic(common.NewCustomError(common.CodeServerBusy))
	}
	common.Success(c, posts)
}
