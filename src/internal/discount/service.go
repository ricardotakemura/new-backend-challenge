package discount

import (
	"math"
)

type DiscountService struct {
	discountModel *DiscountModel
}

var _discountService *DiscountService

func GetDiscountService() *DiscountService {
	if _discountService == nil {
		_discountService = NewDiscountService()
	}
	return _discountService
}

func NewDiscountService() *DiscountService {
	return &DiscountService{discountModel: NewDiscountModel()}
}

func (service DiscountService) CalculateDiscount(productId int32, amount int32) int32 {
	discount := (*service.discountModel).GetDiscount(productId)
	return int32(math.Round(float64(amount) * float64(discount) / 100))
}
