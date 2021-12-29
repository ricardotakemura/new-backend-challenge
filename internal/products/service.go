package products

type ProductService struct {
	productModel *ProductModel
}

func NewProductService() *ProductService {
	return &ProductService{productModel: NewProductModel()}
}

func (service ProductService) List() []Product {
	return (*service.productModel).list()
}
