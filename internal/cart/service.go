package cart

import (
	"errors"
	"new-backend-challenge/internal/product"
	"os"
	"time"
)

type CartService struct {
	productService *product.ProductService
}

func NewCartService() *CartService {
	return &CartService{productService: product.NewProductService()}
}

func (service CartService) CreateCart(request CartRequest) (*Cart, error) {
	cart := Cart{TotalAmount: 0, TotalAmountWithDiscount: 0, TotalDiscount: 0, Items: []CartItem{}}
	for _, item := range request.Items {
		product, err := service.productService.GetById(item.ID)
		if err != nil {
			return nil, err
		}
		if item.Quantity < 1 {
			return nil, errors.New("invalid_quantity")
		}
		totalAmount := (*product).Amount * uint64(item.Quantity)
		cart.TotalAmount += totalAmount
		cart.Items = append(cart.Items, CartItem{ID: (*product).ID,
			Quantity:    item.Quantity,
			UnitAmount:  (*product).Amount,
			TotalAmount: totalAmount,
			Discount:    0,
			IsGift:      false,
		})
	}
	if service.isBlackFriday() && !service.hasGift(cart) {
		gift := service.productService.GetGift()
		cart.Items = append(cart.Items, CartItem{ID: (*gift).ID,
			Quantity:    1,
			UnitAmount:  0,
			TotalAmount: 0,
			Discount:    0,
			IsGift:      true,
		})
	}
	return &cart, nil
}

func (service CartService) isBlackFriday() bool {
	now := time.Now()
	blackFridayDay := os.Getenv("BLACK_FRIDAY_DAY")
	return now.Format("01-02") == blackFridayDay
}

func (service CartService) hasGift(cart Cart) bool {
	for _, item := range cart.Items {
		if item.IsGift {
			return true
		}
	}
	return false
}
