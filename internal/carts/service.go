package carts

type CartService struct {
	model *CartModel
}

func NewCartService() *CartService {
	return &CartService{model: NewCartModel()}
}
