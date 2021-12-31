package product

type Product struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      uint64 `json:"amount"`
	IsGift      bool   `json:"is_gift"`
}
