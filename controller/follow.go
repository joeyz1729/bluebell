package controller

import (
	"strconv"

	"github.com/YiZou89/bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FollowHandler(c *gin.Context) {
	// TODO, to_user_id, action_type
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	toUidStr := c.Query("to_user_id")
	actionStr := c.Query("action_type")
	toUid, err := strconv.ParseUint(toUidStr, 10, 64)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	action, err := strconv.ParseInt(actionStr, 10, 64)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	err = logic.Follow(uid, toUid, action)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, "")

}

func FollowerListHandler(c *gin.Context) {
	// TODO, token 解析uid即可
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	users, err := logic.GetFollowerList(uid)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, users)
}

func FollowingListHandler(c *gin.Context) {

	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	users, err := logic.GetFollowingList(uid)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, users)
}

func FriendListHandler(c *gin.Context) {
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	users, err := logic.GetFriendList(uid)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, users)
}
