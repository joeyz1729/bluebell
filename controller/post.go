package controller

import (
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
	if err != nil {
		zap.L().Error("invalid params", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	post, err := logic.GetPostDetailById(pid)
	if err != nil {
		zap.L().Error("GetPostDetailById() err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, post)
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

func getPageInfo(c *gin.Context) (page, size uint64) {
	pageStr := c.Query("page_num")
	sizeStr := c.Query("page_size")
	var err error
	page, err = strconv.ParseUint(pageStr, 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		size = 10
	}
	return
}
