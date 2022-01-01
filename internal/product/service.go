package product

import (
	"errors"
	"math/rand"
	"strconv"
)

type ProductService struct {
	productModel *ProductModel
}

func NewProductService() *ProductService {
	return &ProductService{productModel: NewProductModel()}
}

func (service ProductService) List() []Product {
	return (*service.productModel).list()
}

func (service ProductService) ListNoGifts() []Product {
	return (*service.productModel).listNoGifts()
}

func (service ProductService) GetById(id int32) (*Product, error) {
	var product = (*service.productModel).getById(id)
	if product == nil || (*product).IsGift {
		return nil, errors.New("product_not_found:" + strconv.FormatUint(uint64(id), 10))
	}
	return product, nil
}

func (service ProductService) GetGift() *Product {
	gifts := service.productModel.listGifts()
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
