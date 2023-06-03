package controller

import (
	"strconv"
	"zouyi/bluebell/dao/mysql"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func CommunityListHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
}

func CommunityDetailHandler(c *gin.Context) {

	communityIdStr := c.Param("id")
	communityId, err := strconv.ParseUint(communityIdStr, 10, 64)
	if err != nil {
		zap.L().Error("[GetCommunityDetailById] [mysql] err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	communityDetail, err := mysql.GetCommunityDetailById(communityId)
	if err != nil {
		zap.L().Error("[GetCommunityDetailById] [mysql] err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityDetail)
}
