package controller

import (
	"errors"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

const ContextUserIdKey = "user_id"
const ContextUsernameKey = "username"

var (
	ErrorUserNotLogin = errors.New("user not login")
)

func GetCurrentUser(c *gin.Context) (userId uint64, err error) {
	var id uint64
	id = 463373410378973185
	c.Set(ContextUserIdKey, id)
	_userId, ok := c.Get(ContextUserIdKey)
	if !ok {
		zap.L().Error("get user id err", zap.Error(err))
		err = ErrorUserNotLogin
		return
	}
	userId, ok = _userId.(uint64)
	if !ok {
		zap.L().Error("user id is not uint64 type", zap.Error(err))
		err = ErrorUserNotLogin
		return
	}
	return
}
