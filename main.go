package main

import (
	"github.com/MoNouri97/go-restapi/src/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	routes.Hello(r)
	// rest api with in memory db
	routes.UsersController(r)
	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
