package main

import (
	"fmt"
	"new-backend-challenge/internal/carts"
	"new-backend-challenge/internal/products"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	productController *products.ProductController
	cartController    *carts.CartController
}

func NewRoutes() *Routes {
	return &Routes{productController: products.NewProductController(), cartController: carts.NewCartController()}
}

func (routes Routes) Routes(server gin.IRouter) {
	server.Use(func(context *gin.Context) {
		var lang = context.Request.Header[`Content-Language`]
		fmt.Println(lang)
	})
	server.GET(`/products`, (*routes.productController).List)
	server.POST(`/carts/checkout`, (*routes.cartController).CreateCart)
}
