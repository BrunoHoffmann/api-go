package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "GET",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "POST",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "list",
			})
		})
	}
}
