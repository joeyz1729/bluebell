package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
)

const ContextUserIdKey = "user_id"
const ContextUsernameKey = "username"

var (
	ErrorUserNotLogin = errors.New("user not login")
)

func GetCurrentUser(c *gin.Context) (userId uint64, err error) {
	_userId, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = _userId.(uint64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}
