package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/test", getTestMessage)

	r.Run()
}

func getTestMessage(c *gin.Context) {
	// Any business logic here would point to something in /pkg
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
