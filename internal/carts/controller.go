package carts

import (
	"net/http"
	"new-backend-challenge/internal/errors"

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
	context.Header("Content-Type", "application/json")
	if err := context.BindJSON(&response); err != nil {
		var errorModel = errors.NewErrorModel("pt-BR")
		context.IndentedJSON(http.StatusUnprocessableEntity, (*errorModel).INVALID_BODY())
		return
	}
	context.IndentedJSON(http.StatusCreated, nil)
}
