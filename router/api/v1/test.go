package v1

import (
	"net/http"
	"standard/internal/e"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, e.SUCCESS.Gin())
}
