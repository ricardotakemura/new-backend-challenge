package carts

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService *CartService
}

func NewCartController() *CartController {
	return &CartController{cartService: NewCartService()}
}

func (controller CartController) CreateCart(context *gin.Context) {
	var response CartRequest
	if err := context.BindJSON(&response); err != nil {
		context.IndentedJSON(http.StatusUnprocessableEntity, err)
		return
	}
}
