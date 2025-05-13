package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})
	r.GET("/about", func(c *gin.Context) {
		c.String(http.StatusOK, "This is about page.")
	})
	r.GET("/greet/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello, %s!", name)
	})
	r.Run() 
}