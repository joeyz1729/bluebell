package controller

import (
	"strconv"

	"github.com/YiZou89/bluebell/model"
	"github.com/go-playground/validator/v10"

	"github.com/YiZou89/bluebell/logic"

	"github.com/YiZou89/bluebell/dao/mysql"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityListHandler 显示所有社区信息
func CommunityListHandler(c *gin.Context) {
	communityList, err := mysql.GetAllCommunityList()
	if err != nil {
		zap.L().Error("GetAllCommunityList() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityJoinListHandler(c *gin.Context) {
	//uid, _, err := GetCurrentUser(c)
	uidStr := c.Query("uid")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	communityList, err := logic.GetCommunityJoinList(uid)
	ResponseSuccess(c, communityList)

}

func UserCommunityListHandler(c *gin.Context) {
	//TODO
}

// CommunityDetailHandler 获取指定社区的详细信息
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64)
	if err != nil {
		zap.L().Error("parse community id err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 根据cid从数据库中提取详细信息
	communityDetail, err := mysql.GetCommunityDetailById(communityId)
	if err != nil {
		zap.L().Error("get community detail err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityDetail)
}

func CommunityJoinHandler(c *gin.Context) {
	var joinForm = new(model.JoinForm)
	if err := c.ShouldBindJSON(&joinForm); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			zap.L().Error("join form validation err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(
				errs.Translate(trans)))
			return
		}
		zap.L().Error("vote form binding err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	uid, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get user id err", zap.Error(err))
		ResponseError(c, CodeInvalidToken)
		return
	}

	err = logic.CommunityJoin(uid, joinForm)
	if err != nil {
		zap.L().Error("join community err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)

}
