package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.POST("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.PUT("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.DELETE("/", func(c *gin.Context) {
		c.String(200, "Hello")
	})

	router.Run(":8080")
}
