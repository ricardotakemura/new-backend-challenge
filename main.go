package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	Routes(router)
	router.Run("localhost:8080")
}
