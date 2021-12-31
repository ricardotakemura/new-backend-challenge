package cart

type CartRequest struct {
	Items []CartItemRequest `json:"products"`
}

type CartItemRequest struct {
	ID       uint `json:"id"`
	Quantity uint `json:"quantity"`
}

type Cart struct {
	TotalAmount             uint64     `json:"total_amount"`
	TotalAmountWithDiscount uint64     `json:"total_amount_with_discount"`
	TotalDiscount           uint64     `json:"total_discount"`
	Items                   []CartItem `json:"products"`
}

type CartItem struct {
	ID          uint   `json:"id"`
	Quantity    uint   `json:"quantity"`
	UnitAmount  uint64 `json:"unit_amount"`
	TotalAmount uint64 `json:"total_amount"`
	Discount    uint64 `json:"discount"`
	IsGift      bool   `json:"is_gift"`
}
