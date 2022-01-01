package product

type Product struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int32  `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
