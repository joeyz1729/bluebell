package controller

import (
	"strconv"

	"github.com/YiZou89/bluebell/logic"
	"github.com/YiZou89/bluebell/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PostHandler 用户登陆后，可以发布帖子
func PostHandler(c *gin.Context) {
	// 解析帖子内容
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("post info bind err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 从token中获取user信息
	userId, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get user by authToken err", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	post.AuthorId = userId

	// 创建帖子并存储
	if err := logic.CreatePost(&post); err != nil {
		zap.L().Error("create post err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, post)
}

// PostListOrderHandler 按照指定顺序获取帖子，如vote数，create_time，update_time等
func PostListOrderHandler(c *gin.Context) {
	PostListForm := &model.PostsForm{
		Page:  1,
		Size:  10,
		Order: model.OrderByTime,
	}
	err := c.ShouldBindQuery(PostListForm)
	if err != nil {
		zap.L().Error("get post list form err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	postList, err := logic.GetPostListInOrder(PostListForm)
	if err != nil {
		zap.L().Error("get post list by order err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)

}

// PostListHandler 分页查找帖子记录，默认为第一页的前十条记录
func PostListHandler(c *gin.Context) {
	pageNum, pageSize := getPageInfo(c)

	postList, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("get post list err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postList)
}

// CommunityPostListHandler 根据社区id获取分页的帖子列表
func CommunityPostListHandler(c *gin.Context) {
	// 获取社区id，分页大小信息
	cidStr := c.Query("cid")
	page, size := getPageInfo(c)
	cid, err := strconv.ParseUint(cidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 查询帖子列表
	postList, err := logic.GetCommunityPostList(cid, page, size)
	if err != nil {
		zap.L().Error("get community post list err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postList)
}

func FavoritePostListHandler(c *gin.Context) {
	// TODO
	uidStr := c.Query("uid")
	page, size := getPageInfo(c)
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	postList, err := logic.GetFavoritePostList(uid, page, size)
	if err != nil {
		zap.L().Error("get user post list err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)
}

// PublishPostListHandler 根据用户id获取分页的帖子列表
func PublishPostListHandler(c *gin.Context) {
	// .../list/?user_id=
	//TODO, 记录帖子投票数，token id是否点赞
	uidStr := c.Query("uid")
	page, size := getPageInfo(c)
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	postList, err := logic.GetPublishPostList(uid, page, size)
	if err != nil {
		zap.L().Error("get user post list err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, postList)

}

// GetPostHandler 获取指定id的帖子详细信息
func GetPostHandler(c *gin.Context) {
	// 获取并解析帖子id
	pidStr := c.Param("pid")
	pid, err := strconv.ParseUint(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}

	// 获取帖子信息
	postDetail, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("get post detail err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postDetail)
}

// getPageInfo 获取分页信息并处理
func getPageInfo(c *gin.Context) (page, size int64) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var err error
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
