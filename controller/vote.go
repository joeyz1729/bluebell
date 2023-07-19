package controller

import (
	"zouyi/bluebell/logic"
	"zouyi/bluebell/model"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func PostVoteHandler(c *gin.Context) {
	// 1. params bind and verification
	var voteForm = new(model.VoteForm)
	if err := c.ShouldBindJSON(&voteForm); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			zap.L().Error("validation err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(
				errs.Translate(trans)))
			return
		}
		zap.L().Error("binding err", zap.Error(err))
		ResponseError(c, CodeInvalidParams)
		return
	}
	userId, _, err := GetCurrentUser(c)
	if err != nil {
		zap.L().Error("invalid user", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	// 2. logic vote
	err = logic.PostVote(userId, voteForm)
	if err != nil {
		zap.L().Error("post vote err", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	// 3. return response
	ResponseSuccess(c, nil)
}
