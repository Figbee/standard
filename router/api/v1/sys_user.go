package v1

import (
	"net/http"
	"standard/app/model"

	"github.com/gin-gonic/gin"
)

// GetUserInfo
// @Summary GetUserInfo 获取用户信息
// @Description 获取用户信息
// @Accept application/json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200
// @Router /v1/base/userinfo [get]
func GetUserInfo(c *gin.Context) {
	claim := c.MustGet("user").(*model.Claim)

	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    claim,
	})

}
