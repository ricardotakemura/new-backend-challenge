package products

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
	fmt.Println("Successfully Opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &model.products)
	return &model
}

func (model ProductModel) list() []Product {
	return model.products
}
