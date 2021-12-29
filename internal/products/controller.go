package products

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService *ProductService
}

func NewProductController() *ProductController {
	return &ProductController{productService: NewProductService()}
}

func (controller ProductController) List(context *gin.Context) {
	context.Header("Content-Type", "application/json")
	context.IndentedJSON(http.StatusOK, (*controller.productService).List())
}
