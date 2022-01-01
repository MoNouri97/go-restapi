package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(r *gin.Engine) {
	hello := r.Group("/hello")
	{
		// Ping test
		hello.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello world!",
			})
		})
		// query string test
		hello.GET("/query", func(c *gin.Context) {
			foo := c.Query("foo")
			bar := c.Query("bar")
			c.JSON(http.StatusOK, gin.H{
				"foo": foo,
				"bar": bar,
			})
		})
		// params test
		hello.GET("/params/:foo/:bar", func(c *gin.Context) {
			foo := c.Param("foo")
			bar := c.Param("bar")
			c.JSON(http.StatusOK, gin.H{
				"foo": foo,
				"bar": bar,
			})
		})
	}

}
