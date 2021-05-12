package main

import (
	"github.com/gin-gonic/gin"
	"klauskie.com/pathfinder/backend/controller"
)

func main() {
	r := gin.Default()

	r.Use(CORSMiddleware)

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "hola",
			})
		})
		api.POST("/calculate", controller.HandleRunResults)
	}

	r.Run("127.0.0.1:8080")
}

func CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, token, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

