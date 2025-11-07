package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
