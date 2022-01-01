package cart

import (
	"errors"
	"new-backend-challenge/internal/config"
	"new-backend-challenge/internal/discount"
	"new-backend-challenge/internal/product"
	"strconv"
	"time"
)

type CartService struct {
	discountService *discount.DiscountService
	productService  *product.ProductService
}

func NewCartService() *CartService {
	return &CartService{productService: product.NewProductService(), discountService: discount.NewDiscountService()}
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
		if service.hasProductInCart(*product, cart) {
			return nil, errors.New("product_already_in_the_cart:" + strconv.FormatInt(int64((*product).ID), 10))
		}
		totalAmount := (*product).Amount * item.Quantity
		discount := service.discountService.CalculateDiscount(product.ID, totalAmount)
		cart.TotalAmount += totalAmount
		cart.TotalDiscount += discount
		cart.TotalAmountWithDiscount += totalAmount - discount
		cart.Items = append(cart.Items, CartItem{ID: (*product).ID,
			Quantity:    item.Quantity,
			UnitAmount:  (*product).Amount,
			TotalAmount: totalAmount,
			Discount:    discount,
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
	blackFridayDay := config.Config()["blackFridayDay"]
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

func (service CartService) hasProductInCart(product product.Product, cart Cart) bool {
	for _, item := range cart.Items {
		if item.ID == product.ID {
			return true
		}
	}
	return false
}
