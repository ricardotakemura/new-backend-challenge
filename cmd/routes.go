package main

import (
	"new-backend-challenge/internal/cart"
	"new-backend-challenge/internal/product"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	productController *product.ProductController
	cartController    *cart.CartController
}

func NewRoutes() *Routes {
	return &Routes{productController: product.NewProductController(), cartController: cart.NewCartController()}
}

func (routes Routes) Routes(server gin.IRouter) {
	server.GET("/products", (*routes.productController).List)
	server.GET("/products/:productid", (*routes.productController).GetById)
	server.POST("/carts/checkout", (*routes.cartController).CreateCart)
}
