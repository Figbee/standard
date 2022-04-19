package middleware

import (
	"net/http"
	"standard/internal/app/model"
	"standard/internal/e"
	"standard/internal/global"

	"github.com/gin-gonic/gin"
)

func PermissionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		claim := c.MustGet("user").(*model.Claim)
		role := claim.RoleName
		enforce, err := global.Casbin.Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			global.Logger.Debugf("casbin enforce err:%v", err)
			c.JSON(http.StatusNotAcceptable, e.AuthFaild.Gin())
			c.Abort()
			return
		}
		if enforce {
			c.Next()
		} else {
			c.JSON(http.StatusForbidden, e.AuthNotAllow.Gin())
			c.Abort()
			return
		}
	}
}
