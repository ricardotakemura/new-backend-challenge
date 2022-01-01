package main

import (
	"new-backend-challenge/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes := NewRoutes()
	(*routes).Routes(server)
	server.Run(config.Config()["port"])
}
