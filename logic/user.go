package logic

import (
	"SnailForum/common"
	"SnailForum/config"
	"SnailForum/dto"
	"SnailForum/model"
	"SnailForum/pkg/jwt"
	"SnailForum/pkg/passwd"
	"SnailForum/pkg/snowflake"
	"SnailForum/vo"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Register(user dto.RegisterDTO) {

	DB := config.GetDB()

	if isUserExist(user.Username) {
		panic(common.NewCustomError(common.CodeUserExist))
	}

	newUser := model.User{
		UserID:   snowflake.GenerateID(),
		Username: user.Username,
		Email:    user.Email,
		Password: passwd.Encode(user.Password),
		Age:      user.Age,
	}

	result := DB.Create(&newUser)
	if result.RowsAffected == 0 {
		panic(errors.New("创建失败"))
	}
}

func Login(user dto.LoginDTO, ctx *gin.Context) vo.Tokens {
	dbUser := FindUserByUsername(user.Username)
	if dbUser.UserID == 0 {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}

	if !passwd.Verify(user.Password, dbUser.Password) {
		panic(common.NewCustomError(common.CodeInvalidPassword))
	}

	accessToken, err := jwt.AccessToken(dbUser.UserID)
	if err != nil {
		panic(err)
	}
	refreshToken, err := jwt.RefreshToken(dbUser.UserID)
	if err != nil {
		panic(err)
	}
	// 将access_token 存入redis中 限制同一用户同一IP 同一时间只能登录一个设备
	// key user:token:user_id:IP value access_token
	config.RDB.Set(context.Background(), common.KeyUserTokenPrefix+strconv.FormatInt(dbUser.UserID, 10)+":"+ctx.RemoteIP(), accessToken, 2*time.Hour)
	return vo.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}

func isUserExist(username string) bool {
	DB := config.GetDB()
	var user model.User
	result := DB.Where(&model.User{Username: username}).Find(&user)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}

func FindUserByUsername(username string) model.User {
	DB := config.GetDB()
	var user model.User
	DB.Where(&model.User{Username: username}).Find(&user)
	return user
}

func GetUserById(userId int64) model.User {
	DB := config.GetDB()
	var user model.User
	DB.Where(&model.User{UserID: userId}).Find(&user)
	return user
}
