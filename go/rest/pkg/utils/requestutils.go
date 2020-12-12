package utils

import (
	"github.com/gin-gonic/gin"
)

// Handle error from http request
func HandleError(c *gin.Context, status int, err string) {
	c.JSON(status, gin.H{
		"code": status,
		"msg":  err,
	})

	c.Abort()
}