package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes := NewRoutes()
	(*routes).Routes(server)
	server.Run("localhost:8080") //config.Config()["port"])
}
