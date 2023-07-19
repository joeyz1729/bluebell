package controller

import (
	"errors"
	"fmt"
	"strings"
	"zouyi/bluebell/dao/mysql"
	"zouyi/bluebell/logic"
	"zouyi/bluebell/model"
	"zouyi/bluebell/pkg/jwt"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

func SignupHandler(c *gin.Context) {
	// 获取请求参数，绑定，参数校验
	sf := new(model.SignupForm)
	if err := c.ShouldBindJSON(&sf); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			zap.L().Error("signup form invalid err", zap.Error(err))
			ResponseError(c, CodeInvalidParams)
			return
		} else {
			zap.L().Error("signup form verification err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(
				errs.Translate(trans)))
			return
		}
	}

	if err := logic.Signup(sf); err != nil {
		zap.L().Error("signup err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return

	}

	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 获取参数，绑定，参数校验
	lf := new(model.LoginForm)
	if err := c.ShouldBindJSON(&lf); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			zap.L().Error("login form invalid err", zap.Error(err))
			ResponseError(c, CodeInvalidParams)
		} else {
			zap.L().Error("login form verification err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		}
		return
	}

	user, err := logic.Login(lf)
	if err != nil {
		zap.L().Error("login logic err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, mysql.ErrorIncorrectPassword) {
			ResponseError(c, CodeInvalidPassword)
			return
		}
		if errors.Is(err, mysql.ErrorQueryUserData) {
			ResponseError(c, CodeServerBusy)
			return
		}
	}

	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserId),
		"user_name":     user.Username,
		"token":         user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

func RefreshTokenHandler(c *gin.Context) {
	// 获取refresh token
	refreshToken := c.Query("refresh_token")

	// 获取access token 并检查格式
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseError(c, CodeNotLogin)
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseError(c, CodeInvalidAuthFormat)
		return
	}

	// refresh token
	aToken, rToken, err := jwt.RefreshToken(parts[1], refreshToken)
	if err != nil {
		zap.L().Error("jwt.refreshToken failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
