package controller

import (
	"fmt"
	"strconv"
	"zouyi/bluebell/logic"
	"zouyi/bluebell/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostHandler(c *gin.Context) {
	// 1. get params and verification
	var post model.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		zap.L().Debug("[PostHandler] bind err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 2. get user id by authToken
	userId, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get current user by authToken err", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}

	post.AuthorId = userId

	// 3. create post and store into database
	if err := logic.CreatePost(&post); err != nil {
		zap.L().Error("create post err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 4. return response
	ResponseSuccess(c, post)
}

func GetPostHandler(c *gin.Context) {
	pidStr := c.Param("pid")
	pid, err := strconv.ParseUint(pidStr, 10, 64)
	fmt.Println(pid)
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	postDetail, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("GetPostDetailById() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postDetail)
}

func PostListOrderHandler(c *gin.Context) {
	// score(vote), create_time, update_time
	// get params from context, order basis
	// GET /api/v1/posts
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
func PostListHandler(c *gin.Context) {

	pageNum, pageSize := getPageInfo(c)

	postList, err := logic.GetPostList(pageNum, pageSize)
	if err != nil {
		zap.L().Error("GetPostList() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, postList)
}

func CommunityPostListHandler(c *gin.Context) {
	// 1. get community id
	// 2. mysql query by community id

}

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
