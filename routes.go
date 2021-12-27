package main

import (
	"new-backend-challenge/products"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.POST("/shopping/cart", products.Post)
}
