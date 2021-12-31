package cart

import (
	"net/http"
	"new-backend-challenge/internal/error"
	"strings"

	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService  *CartService
	errorService *error.ErrorService
}

func NewCartController() *CartController {
	return &CartController{cartService: NewCartService(), errorService: error.NewErrorService()}
}

func (controller CartController) CreateCart(context *gin.Context) {
	var lang = `en`
	if len(context.Request.Header["Content-Language"]) > 0 {
		lang = context.Request.Header["Content-Language"][0]
	}
	var request = CartRequest{}
	context.Header("Content-Type", "application/json")
	if context.BindJSON(&request) != nil {
		context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.errorService).INVALID_BODY(lang))
		return
	}
	var cart, err = controller.cartService.CreateCart(request)
	if err != nil {
		if strings.Contains(err.Error(), "product_not_found:") {
			productId := strings.Replace(err.Error(), "product_not_found:", "", -1)
			context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.errorService).PRODUCT_NOT_FOUND(lang, productId))
			return
		}
		if strings.Contains(err.Error(), "invalid_quantity") {
			context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.errorService).INVALID_QUANTITY(lang))
			return
		}
		context.IndentedJSON(http.StatusUnprocessableEntity, (*controller.errorService).GENERIC_ERROR(lang))
		return
	}
	context.IndentedJSON(http.StatusCreated, *cart)
}
