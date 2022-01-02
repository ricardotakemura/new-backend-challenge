package cart

import (
	"net/http"
	"new-backend-challenge/internal/error"
	"strings"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	CartService  *CartService
	ErrorService *error.ErrorService
}

func NewCartController() *CartController {
	return &CartController{CartService: GetCartService(), ErrorService: error.GetErrorService()}
}

func (controller CartController) CreateCart(context *gin.Context) {
	var lang = `en`
	if len(context.Request.Header["Content-Language"]) > 0 {
		lang = context.Request.Header["Content-Language"][0]
	}
	var request = CartRequest{}
	context.Header("Content-Type", "application/json")
	if context.BindJSON(&request) != nil {
		context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.ErrorService).INVALID_BODY(lang))
		return
	}
	var cart, err = controller.CartService.CreateCart(request)
	if err != nil {
		if strings.Contains(err.Error(), "product_not_found:") {
			productId := strings.Replace(err.Error(), "product_not_found:", "", -1)
			context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.ErrorService).PRODUCT_NOT_FOUND(lang, productId))
			return
		}
		if strings.Contains(err.Error(), "invalid_quantity") {
			context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.ErrorService).INVALID_QUANTITY(lang))
			return
		}
		if strings.Contains(err.Error(), "product_already_in_the_cart") {
			productId := strings.Replace(err.Error(), "product_already_in_the_cart:", "", -1)
			context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.ErrorService).PRODUCT_ALREADY_IN_THE_CART(lang, productId))
			return
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.ErrorService).GENERIC_ERROR(lang))
		return
	}
	context.IndentedJSON(http.StatusCreated, *cart)
}
