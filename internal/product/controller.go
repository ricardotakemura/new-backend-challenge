package product

import (
	"net/http"
	"new-backend-challenge/internal/error"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *ProductService
	errorService   *error.ErrorService
}

func NewProductController() *ProductController {
	return &ProductController{productService: NewProductService(), errorService: error.NewErrorService()}
}

func (controller ProductController) List(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, (*controller.productService).ListNoGifts())
}

func (controller ProductController) GetById(context *gin.Context) {
	lang := "en"
	if len(context.Request.Header["Language"]) > 0 {
		lang = context.Request.Header["Language"][0]
	}
	context.Header("Content-Type", "application/json")
	productIdAsString := context.Param("productid")
	productId, err := strconv.ParseUint(productIdAsString, 10, 16)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, (*controller.errorService).PRODUCT_NOT_FOUND(lang, productIdAsString))
		return
	}
	product, err := (*controller.productService).GetById(uint(productId))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, (*controller.errorService).PRODUCT_NOT_FOUND(lang, productIdAsString))
		return
	}
	context.IndentedJSON(http.StatusOK, *product)
}
