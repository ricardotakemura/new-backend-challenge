package carts

type CartModel struct {
	carts []Cart
}

func NewCartModel() *CartModel {
	return &CartModel{carts: nil}
}
