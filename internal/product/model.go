package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ProductModel struct {
	products []Product
}

func NewProductModel() *ProductModel {
	var model = ProductModel{
		products: []Product{},
	}
	jsonFile, err := os.Open("../static/products.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &model.products)
	return &model
}

func (model ProductModel) list() []Product {
	return model.products
}

func (model ProductModel) listNoGifts() []Product {
	filter := []Product{}
	for _, p := range model.products {
		if !p.IsGift {
			filter = append(filter, p)
		}
	}
	return filter
}

func (model ProductModel) listGifts() []Product {
	filter := []Product{}
	for _, p := range model.products {
		if p.IsGift {
			filter = append(filter, p)
		}
	}
	return filter
}

func (model ProductModel) getById(id uint) *Product {
	for _, p := range model.products {
		if p.ID == id {
			return &p
		}
	}
	return nil
}
