package controller

import (
	"errors"
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
	// 1. get params, params verification
	sf := new(model.SignupForm)
	if err := c.ShouldBindJSON(&sf); err != nil {

		if errs, ok := err.(validator.ValidationErrors); !ok {
			zap.L().Error("[signup] [params] invalid err", zap.Error(err))
			ResponseError(c, CodeInvalidParams)
			return
		} else {
			zap.L().Error("[signup] [params] verification err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(
				errs.Translate(trans)))
			return
		}
	}

	// 2. business process
	if err := logic.Signup(sf); err != nil {
		zap.L().Error("[signup] [logic] err", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		//c.JSON(http.StatusOK, gin.H{
		ResponseError(c, CodeServerBusy)
		return

	}

	// 3. return response
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	// 1. get params and params verification
	lf := new(model.LoginForm)
	if err := c.ShouldBindJSON(&lf); err != nil {
		if errs, ok := err.(validator.ValidationErrors); !ok {
			zap.L().Error("[login] [params] invalid err", zap.Error(err))
			ResponseError(c, CodeInvalidParams)
		} else {
			zap.L().Error("[login] [params] verification err", zap.Error(errs))
			ResponseErrorWithMsg(c, CodeInvalidParams, removeTopStruct(errs.Translate(trans)))
		}
		return
	}
	// 3. logic process
	user, err := logic.Login(lf)
	if err != nil {
		zap.L().Error("[login] [logic] err", zap.Error(err))
		// ErrorUserNotExist
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		// ErrorIncorrectPassword
		if errors.Is(err, mysql.ErrorIncorrectPassword) {
			ResponseError(c, CodeInvalidToken)
		}
		// ErrorQueryUserData
		if errors.Is(err, mysql.ErrorQueryUserData) {
			ResponseError(c, CodeServerBusy)
		}
	}
	// 4. return response

	ResponseSuccess(c, user)
}

func RefreshTokenHandler(c *gin.Context) {
	refreshToken := c.Query("refresh_token")

	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseError(c, CodeNotLogin)
		c.Abort()
		return
	}
	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseError(c, CodeInvalidAuthFormat)
		c.Abort()
		return
	}

	aToken, rToken, err := jwt.RefreshToken(parts[1], refreshToken)
	if err != nil {
		zap.L().Error("jwt.refreshToken failed", zap.Error(err))
		c.Abort()
		return
	}
	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
