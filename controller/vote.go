package controller

import (
	"zouyi/bluebell/logic"
	"zouyi/bluebell/model"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// PostVoteHandler 帖子投票功能
func PostVoteHandler(c *gin.Context) {
	// 获取参数并绑定
	var voteForm = new(model.VoteForm)
	if err := c.ShouldBindJSON(&voteForm); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			zap.L().Error("vote form validation err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(
				errs.Translate(trans)))
			return
		}
		zap.L().Error("vote form binding err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	// 从authToken中获取用户信息
	userId, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("get user info err", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}

	// 记录投票信息
	err = logic.PostVote(userId, voteForm)
	if err != nil {
		zap.L().Error("post vote err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
