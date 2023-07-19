package router

import (
	"net/http"
	"zouyi/bluebell/controller"
	"zouyi/bluebell/logger"
	"zouyi/bluebell/middleware"

	"go.uber.org/zap"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// 自定制日志记录和恢复
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middleware.RateLimitMiddleware(time.Second*2, 1))

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	pprof.Register(r)

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignupHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)
	v1.GET("/community", controller.CommunityListHandler)
	v1.GET("community/:id", controller.CommunityDetailHandler)
	v1.GET("/posts", controller.PostListHandler)
	v1.GET("/posts/:cid", controller.CommunityPostListHandler)
	v1.GET("/post/:pid", controller.GetPostHandler)
	v1.GET("/posts/order", controller.PostListOrderHandler)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/ping", func(c *gin.Context) {
			userId := c.MustGet("user_id")
			controller.ResponseSuccess(c, gin.H{"user_id": userId})
		})
		v1.POST("/post", controller.PostHandler)
		v1.POST("/vote", controller.PostVoteHandler)

	}
	zap.L().Info("[router] init success")
	return r
}
