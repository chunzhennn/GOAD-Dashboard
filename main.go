package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	router.POST("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "flag{fake_flag}",
		})
	})

	router.Run()
}
