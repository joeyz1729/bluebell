package middleware

import (
	"strings"
	"zouyi/bluebell/controller"
	"zouyi/bluebell/pkg/jwt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 检查请求头中的token是否合法，并解析出token中的user信息
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 获取参数
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			zap.L().Error("jwt auth middleware err, not login")
			controller.ResponseError(c, controller.CodeNotLogin)
			c.Abort()
			return
		}
		// 验证参数
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Error("jwt auth middleware err, invalid auth format")
			controller.ResponseError(c, controller.CodeInvalidAuthFormat)
			c.Abort()
			return
		}

		// 解析access token
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Error("jwt auth middleware err, invalid token", zap.Error(err))
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// jwt token确认正确并解析后
		// 将userid和username添加到context中方便后续功能
		c.Set(controller.CtxUserIdKey, claims.UserId)
		c.Set(controller.CtxUsernameKey, claims.Username)
		c.Next()
	}
}
