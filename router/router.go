package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "this is a test route\n")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	r.GET("/greet", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Gin ToDo API!\n")
	})

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "Server is up and running"})
	})

	r.GET("/extra", func(c *gin.Context) {
		c.String(http.StatusOK, "this is an extra route\n")
	})

	return r
}
