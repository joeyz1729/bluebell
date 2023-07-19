package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const CtxUserIdKey = "user_id"
const CtxUsernameKey = "username"

var (
	ErrorUserNotLogin = errors.New("user not login")
)

// GetCurrentUser jwt鉴权后将用户信息放入context中，通过context获取用户信息
func GetCurrentUser(c *gin.Context) (userId uint64, username string, err error) {
	// 从context中获取user id
	var ok bool
	id, ok := c.Get(CtxUserIdKey)
	name, ok := c.Get(CtxUsernameKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	userId, ok = id.(uint64)
	username, ok = name.(string)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
