package controller

import (
	"database/sql"
	"strconv"

	"github.com/YiZou89/bluebell/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FollowHandler(c *gin.Context) {
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
	var attitude bool
	if action == int64(1) {
		attitude = true
	} else if action == int64(2) {
		attitude = false
	} else {
		zap.L().Error("invalid action type")
		ResponseError(c, CodeInvalidParams)
		return
	}

	err = logic.Follow(uid, toUid, attitude)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, "")

}

func FollowerListHandler(c *gin.Context) {
	// jwt token 中拿到uid
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 尝试获取query中参数, 如果正确就查询需要的user id
	uidStr := c.Query("user_id")
	if uidStr != "" {
		userId, err := strconv.ParseUint(uidStr, 10, 64)
		if err == nil {
			uid = userId
		}
	}

	followers, err := logic.GetFollowers(uid)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, followers)

}

func FollowingListHandler(c *gin.Context) {
	// jwt token 中拿到uid
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 尝试获取query中参数, 如果正确就查询需要的user id
	uidStr := c.Query("user_id")
	if uidStr != "" {
		userId, err := strconv.ParseUint(uidStr, 10, 64)
		if err == nil {
			uid = userId
		}
	}

	followers, err := logic.GetFollowings(uid)
	if err != nil {
		if err == sql.ErrNoRows {
			ResponseSuccess(c, nil)
			return
		}
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, followers)
}

func FriendListHandler(c *gin.Context) {
	// jwt token 中拿到uid
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 尝试获取query中参数, 如果正确就查询需要的user id
	uidStr := c.Query("user_id")
	if uidStr != "" {
		userId, err := strconv.ParseUint(uidStr, 10, 64)
		if err == nil {
			uid = userId
		}
	}

	followers, err := logic.GetFriends(uid)
	if err != nil {
		zap.L().Error("", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	ResponseSuccess(c, followers)
}
