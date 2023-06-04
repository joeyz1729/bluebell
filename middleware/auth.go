package middleware

import (
	"strings"
	"zouyi/bluebell/controller"
	"zouyi/bluebell/pkg/jwt"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware check authorization,
// set userId
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			zap.L().Error("code not login err")
			controller.ResponseError(c, controller.CodeNotLogin)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			zap.L().Error("invalid auth format")
			controller.ResponseError(c, controller.CodeInvalidAuthFormat)
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(parts[1])
		if err != nil {
			zap.L().Error("invalid token", zap.Error(err))
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}
		c.Set(controller.ContextUserIdKey, claims.UserId)
		c.Next()
	}
}
