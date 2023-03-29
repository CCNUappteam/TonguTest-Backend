package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"tongue/handler"
	"tongue/handler/auth"
	"tongue/handler/sd"
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
	}

	// user 模块
	//userRouter := g.Group("api/v1/user")
	//{
	//	userRouter.POST("/login", user.Login)
	//
	//	userRouter.PUT("/password", normalRequired, user.ChangePassword)
	//
	//	userRouter.PUT("", normalRequired, user.UpdateInfo)
	//
	//	userRouter.GET("/info", normalRequired, user.GetInfo)
	//
	//	userRouter.GET("/profile/:email", user.GetProfile)
	//
	//	// userRouter.GET("/list", user.List)
	//
	//	userRouter.GET("/qiniu_token", user.GetQiniuToken)
	//
	//	userRouter.PUT("/role", normalRequired, user.SetRole)
	//
	//}

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
