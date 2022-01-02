package product

import (
	"errors"
	"math/rand"
	"strconv"
)

type ProductService struct {
	ProductModel IProductModel
}

var _productService *ProductService

func GetProductService() *ProductService {
	if _productService == nil {
		_productService = NewProductService()
	}
	return _productService
}

func NewProductService() *ProductService {
	return &ProductService{ProductModel: NewProductModel()}
}

func (service ProductService) List() []Product {
	return service.ProductModel.List()
}

func (service ProductService) ListNoGifts() []Product {
	return service.ProductModel.ListNoGifts()
}

func (service ProductService) GetById(id int32) (*Product, error) {
	var product = service.ProductModel.GetById(id)
	if product == nil || (*product).IsGift {
		return nil, errors.New("product_not_found:" + strconv.FormatUint(uint64(id), 10))
	}
	return product, nil
}

func (service ProductService) GetGift() *Product {
	gifts := service.ProductModel.ListGifts()
	size := len(gifts)
	if size == 0 {
		return nil
	}
	var n = rand.Intn(size)
	if size == 1 || n < 0 {
		return &gifts[0]
	}
	return &gifts[n]
}
