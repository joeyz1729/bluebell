package middleware

import (
	"strings"
	"zouyi/bluebell/controller"
	"zouyi/bluebell/pkg/jwt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 检查请求头中的token是否合法
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			zap.L().Error("[jwt auth middleware] not login")
			controller.ResponseError(c, controller.CodeNotLogin)
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Error("[jwt auth middleware] invalid auth format")
			controller.ResponseError(c, controller.CodeInvalidAuthFormat)
			c.Abort()
			return
		}

		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Error("[jwt auth middleware] invalid token", zap.Error(err))
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// jwt token确认正确并解析后，
		// 将userid和username添加到context中方便后续功能
		c.Set(controller.CtxUserIdKey, claims.UserId)
		c.Set(controller.CtxUsernameKey, claims.Username)
		c.Next()
	}
}
