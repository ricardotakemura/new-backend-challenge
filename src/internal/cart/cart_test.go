package cart

import (
	"new-backend-challenge/internal/discount"
	"new-backend-challenge/internal/product"
	"os"
	"strings"
	"testing"
	"time"
)

type MockDiscountModel struct {
}

func (model MockDiscountModel) GetDiscount(productId int32) float32 {
	return float32(productId)
}

func Test_CartService_CreateCart_Ok(t *testing.T) {
	request := CartRequest{Items: []CartItemRequest{{ID: 1, Quantity: 10}}}
	discountModel := &MockDiscountModel{}
	productModel := &product.ProductModel{Products: []product.Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	cartService := CartService{
		ProductService:  &product.ProductService{ProductModel: productModel},
		DiscountService: &discount.DiscountService{DiscountModel: discountModel}}
	cart, error := cartService.CreateCart(request)
	if error != nil {
		t.Fatal(error)
		return
	}
	if len(cart.Items) != 1 {
		t.Errorf("len(cart.Items): expected %d obtained %d", 1, len(cart.Items))
	}
	if cart.TotalAmount != 1000 {
		t.Errorf("cart.TotalAmount: expected %d obtained %d", 1000, cart.TotalAmount)
	}
	if cart.TotalAmountWithDiscount != 990 {
		t.Errorf("cart.TotalAmountWithDiscount: expected %d obtained %d", 990, cart.TotalAmountWithDiscount)
	}
	if cart.TotalDiscount != 10 {
		t.Errorf("cart.TotalDiscount: expected %d obtained %d", 10, cart.TotalDiscount)
	}
}

func Test_CartService_CreateCart_ProductNotFound(t *testing.T) {
	request := CartRequest{Items: []CartItemRequest{{ID: 1, Quantity: 10}}}
	discountModel := &MockDiscountModel{}
	productModel := &product.ProductModel{Products: []product.Product{}}
	cartService := CartService{
		ProductService:  &product.ProductService{ProductModel: productModel},
		DiscountService: &discount.DiscountService{DiscountModel: discountModel}}
	_, error := cartService.CreateCart(request)
	if error == nil || !strings.Contains(error.Error(), "product_not_found") {
		t.Errorf("cartService.CreateCart: excepted \"product_not_found\" obtained nil")
	}
}

func Test_CartService_CreateCart_Invalid_Quantity(t *testing.T) {
	request := CartRequest{Items: []CartItemRequest{{ID: 1, Quantity: 0}}}
	discountModel := &MockDiscountModel{}
	productModel := &product.ProductModel{Products: []product.Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	cartService := CartService{
		ProductService:  &product.ProductService{ProductModel: productModel},
		DiscountService: &discount.DiscountService{DiscountModel: discountModel}}
	_, error := cartService.CreateCart(request)
	if error == nil || !strings.Contains(error.Error(), "invalid_quantity") {
		t.Errorf("cartService.CreateCart: excepted \"invalid_quantity\" obtained nil")
	}
}

func Test_CartService_CreateCart_Duplicated_Product(t *testing.T) {
	request := CartRequest{Items: []CartItemRequest{{ID: 1, Quantity: 1}, {ID: 1, Quantity: 1}}}
	discountModel := &MockDiscountModel{}
	productModel := &product.ProductModel{Products: []product.Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	cartService := CartService{
		ProductService:  &product.ProductService{ProductModel: productModel},
		DiscountService: &discount.DiscountService{DiscountModel: discountModel}}
	_, error := cartService.CreateCart(request)
	if error == nil || !strings.Contains(error.Error(), "product_already_in_the_cart") {
		t.Errorf("cartService.CreateCart: excepted \"product_already_in_the_cart\" obtained nil")
	}
}

func Test_CartService_CreateCart_WithGift(t *testing.T) {
	os.Setenv("BLACK_FRIDAY_DAY", time.Now().Format("01-02"))
	request := CartRequest{Items: []CartItemRequest{{ID: 1, Quantity: 10}}}
	discountModel := &MockDiscountModel{}
	productModel := &product.ProductModel{Products: []product.Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	cartService := CartService{
		ProductService:  &product.ProductService{ProductModel: productModel},
		DiscountService: &discount.DiscountService{DiscountModel: discountModel}}
	cart, error := cartService.CreateCart(request)
	if error != nil {
		t.Fatal(error)
		return
	}
	if len(cart.Items) != 2 {
		t.Errorf("len(cart.Items): expected %d obtained %d", 2, len(cart.Items))
	}
	if cart.TotalAmount != 1000 {
		t.Errorf("cart.TotalAmount: expected %d obtained %d", 1000, cart.TotalAmount)
	}
	if cart.TotalAmountWithDiscount != 990 {
		t.Errorf("cart.TotalAmountWithDiscount: expected %d obtained %d", 990, cart.TotalAmountWithDiscount)
	}
	if cart.TotalDiscount != 10 {
		t.Errorf("cart.TotalDiscount: expected %d obtained %d", 10, cart.TotalDiscount)
	}
}
