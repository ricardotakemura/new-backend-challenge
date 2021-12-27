package main

import (
	"new-backend-challenge/products"

	"github.com/gin-gonic/gin"
)

func Routes(router gin.IRouter) {
	router.POST("/shopping/cart", products.Post)
}
