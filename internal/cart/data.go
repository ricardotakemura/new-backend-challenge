package cart

type CartRequest struct {
	Items []CartItemRequest `json:"products"`
}

type CartItemRequest struct {
	ID       int32 `json:"id"`
	Quantity int32 `json:"quantity"`
}

type Cart struct {
	TotalAmount             int32      `json:"total_amount"`
	TotalAmountWithDiscount int32      `json:"total_amount_with_discount"`
	TotalDiscount           int32      `json:"total_discount"`
	Items                   []CartItem `json:"products"`
}

type CartItem struct {
	ID          int32 `json:"id"`
	Quantity    int32 `json:"quantity"`
	UnitAmount  int32 `json:"unit_amount"`
	TotalAmount int32 `json:"total_amount"`
	Discount    int32 `json:"discount"`
	IsGift      bool  `json:"is_gift"`
}
