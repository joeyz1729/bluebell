package controller

import (
	"strconv"

	"github.com/YiZou89/bluebell/dao/mysql"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// CommunityListHandler 显示所有社区信息
func CommunityListHandler(c *gin.Context) {
	communityList, err := mysql.GetCommunityList()
	if err != nil {
		zap.L().Error("GetCommunityList() error", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, communityList)
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
