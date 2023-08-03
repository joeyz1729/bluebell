package router

import (
	"net/http"

	"github.com/YiZou89/bluebell/logic"

	"github.com/YiZou89/bluebell/middleware"

	"github.com/YiZou89/bluebell/controller"
	"github.com/YiZou89/bluebell/logger"
	httptransport "github.com/go-kit/kit/transport/http"
	"go.uber.org/zap"

	"github.com/gin-contrib/pprof"

	"github.com/gin-gonic/gin"
)

func Setup(mode string) *gin.Engine {
	zap.L().Debug("setup router")
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

	us := logic.NewUserService()
	getUserHandler := httptransport.NewServer(
		makeGetUserEndpoint(us),
		decodeUserRequest,
		encodeUserResponse,
	)
	r.GET("/username", func(c *gin.Context) {
		getUserHandler.ServeHTTP(c.Writer, c.Request)
	})

	pprof.Register(r)
	v2 := r.Group("/api/v2")

	// 用户
	v2.POST("/signup", controller.SignupHandler)
	v2.POST("/login", controller.LoginHandler)
	v2.GET("/refresh_token", controller.RefreshTokenHandler)
	v2.GET("/userinfo/", middleware.JWTAuthMiddleware(), controller.GetUserInfoHandler) // 获取用户信息

	// 社区
	v2.GET("/communities", controller.CommunityListHandler)
	v2.GET("/join/list/", middleware.JWTAuthMiddleware(), controller.CommunityJoinListHandler)
	v2.GET("community/detail/:id", controller.CommunityDetailHandler)
	v2.POST("/member/action/", middleware.JWTAuthMiddleware(), controller.CommunityJoinHandler)

	// 帖子
	v2.GET("/posts", controller.PostListHandler)
	v2.GET("/community/list/", controller.CommunityPostListHandler)
	v2.GET("/publish/list/", middleware.JWTAuthMiddleware(), controller.PublishPostListHandler) // 获取userid 发布的所有帖子
	v2.GET("/post/:pid", controller.GetPostHandler)
	v2.GET("/posts/order", controller.PostListOrderHandler)
	v2.POST("/post", middleware.JWTAuthMiddleware(), controller.PostHandler)
	v2.POST("/vote", middleware.JWTAuthMiddleware(), controller.PostVoteHandler)

	// 关注
	v2.POST("/relation/action/", middleware.JWTAuthMiddleware(), controller.FollowHandler)
	v2.GET("/relation/following/list/", middleware.JWTAuthMiddleware(), controller.FollowingListHandler)
	v2.GET("/relation/follower/list/", middleware.JWTAuthMiddleware(), controller.FollowerListHandler)
	v2.GET("/relation/friend/list/", middleware.JWTAuthMiddleware(), controller.FriendListHandler)

	// 私信
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
