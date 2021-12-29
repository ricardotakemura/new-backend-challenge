package carts

type CartRequest struct {
	Items []CartItemRequest `json:"products"`
}

type CartItemRequest struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type Cart struct {
	TotalAmount             float64    `json:"total_amount"`
	TotalAmountWithDiscount float64    `json:"total_amount_with_discount"`
	TotalDiscount           float64    `json:"total_discount"`
	Items                   []CartItem `json:"products"`
}

type CartItem struct {
	ID          uint    `json:"id"`
	Quantity    uint    `json:"quantity"`
	UnitAmount  float64 `json:"unit_amount"`
	TotalAmount float64 `json:"total_amount"`
	Discount    float64 `json:"discount"`
	IsGift      bool    `json:"is_gift"`
}
