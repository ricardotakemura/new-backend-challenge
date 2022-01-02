package discount

import (
	"testing"
)

type MockDiscountModel struct {
}

func (model MockDiscountModel) GetDiscount(productId int32) float32 {
	return float32(productId)
}

func Test_DiscountService_CalculateDiscount_Ok(t *testing.T) {
	discountService := DiscountService{DiscountModel: &MockDiscountModel{}}
	discount := discountService.CalculateDiscount(1, 200)
	if discount != 2 {
		t.Errorf("discountService.CalculateDiscount: expected %d obtained %d", 2, discount)
	}
}

func Test_DiscountService_CalculateDiscount_Amount_Negative(t *testing.T) {
	discountService := DiscountService{DiscountModel: &MockDiscountModel{}}
	discount := discountService.CalculateDiscount(1, -200)
	if discount != 0 {
		t.Errorf("discountService.CalculateDiscount: expected %d obtained %d", 0, discount)
	}
}

func Test_DiscountService_CalculateDiscount_Discount_Negative(t *testing.T) {
	discountService := DiscountService{DiscountModel: &MockDiscountModel{}}
	discount := discountService.CalculateDiscount(-1, 100)
	if discount != 0 {
		t.Errorf("discountService.CalculateDiscount: expected %d obtained %d", 0, discount)
	}
}
