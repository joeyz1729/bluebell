package router

import (
	"net/http"

	"github.com/YiZou89/bluebell/middleware"

	"github.com/YiZou89/bluebell/controller"
	"github.com/YiZou89/bluebell/logger"
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

	v2 := r.Group("/api/v2")

	// 社区
	v2.GET("/community", controller.CommunityListHandler)
	v2.GET("community/:id", controller.CommunityDetailHandler)

	// v2.POST("/community/publish", controller.TODO)
	v2.GET("/community/list/:userid", controller.TODO)                          // 获取userid关注的community列表
	v2.POST("/community/join", middleware.JWTAuthMiddleware(), controller.TODO) // userid关注community

	// 用户
	v2.POST("/signup", controller.SignupHandler)
	v2.POST("/login", controller.LoginHandler)
	v2.GET("/userinfo/uid", middleware.JWTAuthMiddleware(), controller.GetUserInfoHandler) // 获取用户信息
	v2.GET("/refresh_token", controller.RefreshTokenHandler)

	v2.GET("/posts/favorite", controller.TODO) // 获取帖子点赞列表
	v2.GET("/posts", controller.PostListHandler)
	v2.GET("/posts/:cid", controller.CommunityPostListHandler)
	v2.GET("/post/:pid", controller.GetPostHandler)
	v2.GET("/posts/order", controller.PostListOrderHandler)
	v2.POST("/post", middleware.JWTAuthMiddleware(), controller.PostHandler)
	v2.POST("/vote", middleware.JWTAuthMiddleware(), controller.PostVoteHandler)

	// 关注或取关， 关注，粉丝，好友列表
	v2.POST("/relation/action/", controller.TODO)
	v2.GET("/relation/follow/list/", controller.TODO)
	v2.GET("/relation/follower/list/", controller.TODO)
	v2.GET("/relation/friend/list/", controller.TODO)

	v2.POST("/message/action/", controller.TODO)
	v2.GET("/message/chat/", controller.TODO)
	v2.GET("/ping", middleware.JWTAuthMiddleware(), func(c *gin.Context) {
		userId := c.MustGet("user_id")
		controller.ResponseSuccess(c, gin.H{"user_id": userId})
	})

	//v2.Use(middleware.JWTAuthMiddleware())
	//{

	//}

	zap.L().Info("[router] init success")
	return r
}
