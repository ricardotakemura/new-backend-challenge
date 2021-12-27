package products

type ProductRequest struct {
	ID       string `json:"id"`
	Quantity uint   `json:"quantity"`
}

type CartRequest struct {
	Products []ProductRequest `json:"products"`
}

type CartResponse struct {
	TotalAmount             float64 `json:"total_amount"`
	TotalAmountWithDiscount float64 `json:"total_amount_with_discount"`
	TotalDiscount           float64 `json:"total_discount"`
}

type ProductResponse struct {
	ID          string  `json:"id"`
	Quantity    uint    `json:"quantity"`
	UnitAmount  float64 `json:"unit_amount"`
	TotalAmount float64 `json:"total_amount"`
	Discount    float64 `json:"discount"`
	IsGift      bool    `json:"is_gift"`
}
