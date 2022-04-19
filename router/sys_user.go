package router

import (
	"standard/app/middleware"
	v1 "standard/router/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) (R gin.IRoutes) {
	router := r.Group("base").Use(authMiddleware.MiddlewareFunc()).Use(middleware.PermissionMiddleWare())
	{
		router.GET("/userinfo", v1.GetUserInfo)
	}
	return router
}
