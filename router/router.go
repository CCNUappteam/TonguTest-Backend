package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"tongue/handler"
	"tongue/handler/auth"
	"tongue/handler/forum"
	"tongue/handler/sd"
	"tongue/handler/test"
	"tongue/handler/user"
	"tongue/handler/xinbang"
	"tongue/pkg/errno"
	"tongue/router/middleware"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		handler.SendError(c, errno.ErrIncorrectAPIRoute, nil, "", "")
	})

	// swagger API doc
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//normalRequired := middleware.AuthMiddleware(constvar.AuthLevelNormal)
	//adminRequired := middleware.AuthMiddleware(constvar.AuthLevelAdmin)
	//superAdminRequired := middleware.AuthMiddleware(constvar.AuthLevelSuperAdmin)

	// auth
	authRouter := g.Group("api/v1/auth")
	{
		authRouter.POST("/code", auth.VerificationCode)
		authRouter.POST("/register", auth.Register)
		authRouter.POST("/login", user.Login)
	}

	// user 模块
	userRouter := g.Group("api/v1/user").Use(middleware.AuthMiddleware())
	{
		userRouter.GET("/info", user.GetInfo)
		userRouter.POST("/avatar", user.UploadAvatar)
		userRouter.POST("/info", user.UpdateInfo)
		userRouter.POST("/card", user.PunchCard)
		userRouter.GET("/card", user.GetCard)
	}

	// forum
	forumRouter := g.Group("api/v1/forum")
	{
		forumRouter.POST("/post", middleware.AuthMiddleware(), forum.PublishPost)
		forumRouter.DELETE("/post", middleware.AuthMiddleware(), forum.DeletePost)
		forumRouter.POST("/image", middleware.AuthMiddleware(), forum.PostImage)
		forumRouter.GET("/posts", forum.GetPosts)
		forumRouter.GET("/myposts", middleware.AuthMiddleware(), forum.GetMyPosts)
		forumRouter.GET("/postimage", forum.GetImage)
		forumRouter.POST("/comment", middleware.AuthMiddleware(), forum.PostComment)
		forumRouter.DELETE("/comment", middleware.AuthMiddleware(), forum.DeleteComment)
		forumRouter.GET("/comment", middleware.AuthMiddleware(), forum.GetComment)
		forumRouter.POST("/like", middleware.AuthMiddleware(), forum.PostLike)
		forumRouter.DELETE("/like", middleware.AuthMiddleware(), forum.DeleteLike)
		forumRouter.GET("/like", forum.GetLikes)
	}

	testRouter := g.Group("api/v1/test").Use(middleware.AuthMiddleware())
	{
		testRouter.POST("/tongue", test.TongueTest)
	}

	rankRouter := g.Group("api/v1/rank")
	{
		rankRouter.POST("", xinbang.GetRank)
	}
	// The health check Handlers
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
