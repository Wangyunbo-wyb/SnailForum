package api

import (
	"SnailForum/common"
	"SnailForum/config"
	"SnailForum/dto"
	"SnailForum/logic"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserInfoByID 用户详情
func GetUserInfoByID(ctx *gin.Context) {
	/*uid := util.StringToInt(ctx.Query("uid"))
	res := service.GetUserInfoByIDService(uid)
	response.HandleResponse(ctx, res)*/
	value, _ := ctx.Get("userId")
	userId, _ := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)

	ctx.JSON(http.StatusOK, gin.H{
		"msg": logic.GetUserById(userId),
	})
}

// Register 用户注册
func Register(ctx *gin.Context) {
	var registerDTO dto.RegisterDTO
	if err := ctx.ShouldBind(&registerDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	logic.Register(registerDTO)

	common.Success(ctx, nil)
}

// Login 用户登录
func Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		config.ValidateError(ctx, err)
		return
	}
	tokens := logic.Login(loginDTO, ctx)
	common.Success(ctx, tokens)
}
