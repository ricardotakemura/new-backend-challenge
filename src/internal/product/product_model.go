package product

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type IProductModel interface {
	List() []Product
	ListNoGifts() []Product
	ListGifts() []Product
	GetById(id int32) *Product
}

type ProductModel struct {
	Products []Product
}

func NewProductModel() *ProductModel {
	var model = ProductModel{
		Products: []Product{},
	}
	jsonFile, err := os.Open("../static/products.json")
	if err != nil {
		log.Fatalf("Failed to open products.json: %s", err.Error())
		return &model
	}
	log.Println("Successfully opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &model.Products)
	return &model
}

func (model ProductModel) List() []Product {
	return model.Products
}

func (model ProductModel) ListNoGifts() []Product {
	filter := []Product{}
	for _, p := range model.Products {
		if !p.IsGift {
			filter = append(filter, p)
		}
	}
	return filter
}

func (model ProductModel) ListGifts() []Product {
	filter := []Product{}
	for _, p := range model.Products {
		if p.IsGift {
			filter = append(filter, p)
		}
	}
	return filter
}

func (model ProductModel) GetById(id int32) *Product {
	for _, p := range model.Products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
