package passwd

import (
	"crypto/sha512"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"go.uber.org/zap"
	"strings"
)

var options = &password.Options{SaltLen: 16, Iterations: 100, KeyLen: 32, HashFunction: sha512.New}

func Encode(rawPassword string) string {
	salt, encodedPwd := password.Encode(rawPassword, options)
	dbPassword := fmt.Sprintf("$sha512$%s$%s", salt, encodedPwd)
	return dbPassword
}

func Verify(rawPassword string, dbPassword string) bool {
	splits := strings.Split(dbPassword, "$")
	if len(splits) != 4 {
		zap.L().Error("数据库密码格式错误")
		return false
	}
	valid := password.Verify(rawPassword, splits[2], splits[3], options)
	return valid
}
