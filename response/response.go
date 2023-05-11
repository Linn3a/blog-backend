package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(c *gin.Context, httpStatus int, code bool, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"state": gin.H{
		"ok":      code,
		"message": msg,
	}, "data": data})
}

// Success 成功
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, true, data, msg)
}

// Fail 失败
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, false, data, msg)
}
