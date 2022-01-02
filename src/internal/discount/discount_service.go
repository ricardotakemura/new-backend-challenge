package discount

import (
	"math"
)

type DiscountService struct {
	DiscountModel IDiscountModel
}

var _discountService *DiscountService

func GetDiscountService() *DiscountService {
	if _discountService == nil {
		_discountService = NewDiscountService()
	}
	return _discountService
}

func NewDiscountService() *DiscountService {
	return &DiscountService{DiscountModel: NewDiscountModel()}
}

func (service DiscountService) CalculateDiscount(productId int32, amount int32) int32 {
	if amount <= 0 {
		return 0
	}
	discount := (service.DiscountModel).GetDiscount(productId)
	if discount <= 0 {
		return 0
	}
	return int32(math.Round(float64(amount) * float64(discount) / 100))
}
