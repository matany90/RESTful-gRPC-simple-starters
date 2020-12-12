package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/middleware"
	v1 "github.com/matany90/RESTful-gRPC-simple-starters/go/rest/pkg/routers/api/v1"
)

// InitRouters will init
// the routing server system using GIN framework.
func InitRouters() *gin.Engine {
	// init gin
	r := gin.New()

	// set middlewares
	// r.Use(middleware.Logger())
	r.Use(gin.Recovery())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})

	// define 404 handler
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"code": 404,
			"msg": "page not found.",
		})
	})

	// group v1 api
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.VerifyAccess())
	{
		apiV1.GET("/matan", v1.GetMatan)
	}

	// return router instance
	return r
}