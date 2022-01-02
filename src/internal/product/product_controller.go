package product

import (
	"net/http"
	"new-backend-challenge/internal/error"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *ProductService
	ErrorService   *error.ErrorService
}

func NewProductController() *ProductController {
	return &ProductController{ProductService: GetProductService(), ErrorService: error.GetErrorService()}
}

func (controller ProductController) List(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, (*controller.ProductService).ListNoGifts())
}

func (controller ProductController) GetById(context *gin.Context) {
	lang := "en"
	if len(context.Request.Header["Content-Language"]) > 0 {
		lang = context.Request.Header["Content-Language"][0]
	}
	context.Header("Content-Type", "application/json")
	productIdAsString := context.Param("productid")
	productId, err := strconv.ParseInt(productIdAsString, 10, 16)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, (*controller.ErrorService).PRODUCT_NOT_FOUND(lang, productIdAsString))
		return
	}
	product, err := (*controller.ProductService).GetById(int32(productId))
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, (*controller.ErrorService).PRODUCT_NOT_FOUND(lang, productIdAsString))
		return
	}
	context.IndentedJSON(http.StatusOK, *product)
}
