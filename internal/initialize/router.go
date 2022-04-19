package initialize

import (
	"standard/app/middleware"
	_ "standard/docs"
	"standard/internal/global"
	"standard/router"
	v1 "standard/router/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//初始化路由
func Routers() *gin.Engine {
	r := gin.Default()
	//gin 使用通用中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Translations())

	authMiddleware, err := middleware.InitAuth()
	if err != nil {
		global.Logger.Panicf("初始化jwt auth中间件失败%v", err)
	}
	r.GET("/ping", v1.Ping)
	versionGroup := r.Group("/v1")
	versionGroup.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//初始化公共路由
	router.InitPublicRouter(versionGroup, authMiddleware)

	router.InitUserRouter(versionGroup, authMiddleware)
	return r
}
