package controller

import (
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const CtxUserIdKey = "user_id"
const CtxUsernameKey = "username"

var (
	ErrorUserNotLogin = errors.New("user not login")
)

// GetCurrentUser 从context中获取uid信息，可用于jwt鉴权后
func GetCurrentUser(c *gin.Context) (userId uint64, err error) {
	uid, ok := c.Get(CtxUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		zap.L().Error("[request] not login", zap.Error(err))
		return
	}

	userId, ok = uid.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		zap.L().Error("[request] invalid user_id", zap.Error(err))
		return
	}
	return
}
