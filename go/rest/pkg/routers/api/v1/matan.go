package v1

import (
	"github.com/gin-gonic/gin"
)
// @Summary Returns "matan" instance
// @Produce  json
// @Success 200 {object} c.JSON
// @Router /api/v1/matan [list]
func GetMatan(c *gin.Context) {
	c.JSON(200, gin.H{"matan": true})
}