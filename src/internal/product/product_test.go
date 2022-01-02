package product

import (
	"testing"
)

func Test_ProductService_List_Ok(t *testing.T) {
	productModel := &ProductModel{Products: []Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	productService := ProductService{ProductModel: productModel}
	products := productService.List()
	if len(products) != 2 {
		t.Errorf("len(products): expected %d obtained %d", 2, len(products))
	}
}

func Test_ProductService_List_Empty_List(t *testing.T) {
	productModel := &ProductModel{Products: []Product{}}
	productService := ProductService{ProductModel: productModel}
	products := productService.List()
	if len(products) != 0 {
		t.Errorf("len(products): expected %d obtained %d", 0, len(products))
	}
}

func Test_ProductService_ListNoGifts_Ok(t *testing.T) {
	productModel := &ProductModel{Products: []Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	productService := ProductService{ProductModel: productModel}
	products := productService.ListNoGifts()
	if len(products) != 1 {
		t.Errorf("len(products): expected %d obtained %d", 1, len(products))
	}
}

func Test_ProductService_GetGift_Ok(t *testing.T) {
	productModel := &ProductModel{Products: []Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: true},
	}}
	productService := ProductService{ProductModel: productModel}
	product := productService.GetGift()
	if product == nil {
		t.Errorf("product: expected product<struts> obtained nill")
		return
	}
	if product.ID != 2 {
		t.Errorf("product.ID: expected %d obtained %d", 2, product.ID)
	}
	if product.Title != "Abacaxi" {
		t.Errorf("product.Title: expected %s obtained %s", "Abacaxi", product.Title)
	}
}

func Test_ProductService_GetGift_No_Gift(t *testing.T) {
	productModel := &ProductModel{Products: []Product{
		{ID: 1, Title: "Banana", Description: "Fruta", Amount: 100, IsGift: false},
		{ID: 2, Title: "Abacaxi", Description: "Fruta", Amount: 200, IsGift: false},
	}}
	productService := ProductService{ProductModel: productModel}
	product := productService.GetGift()
	if product != nil {
		t.Errorf("product: expected nil obtained %d", product.ID)
		return
	}

}
