package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes := NewRoutes()
	(*routes).Routes(server)
	server.Run(":8080")
}
