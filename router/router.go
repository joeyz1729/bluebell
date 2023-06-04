package router

import (
	"zouyi/bluebell/controller"
	"zouyi/bluebell/logger"
	"zouyi/bluebell/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.DebugMode)

	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1.POST("/signup", controller.SignupHandler)
	v1.POST("/login", controller.LoginHandler)
	v1.GET("/refresh_token", controller.RefreshTokenHandler)

	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/ping", func(c *gin.Context) {
			userId := c.MustGet("user_id")
			controller.ResponseSuccess(c, gin.H{"user_id": userId})
		})

		v1.GET("/community", controller.CommunityListHandler)
		v1.GET("community/:id", controller.CommunityDetailHandler)

		v1.POST("/post", controller.PostHandler)
		v1.GET("/posts", controller.PostListHandler)
		v1.GET("/post/:pid", controller.GetPostHandler)

		v1.POST("/vote", controller.PostVoteHandler)

	}

	return r
}
