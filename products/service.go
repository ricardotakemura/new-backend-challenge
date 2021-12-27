package products

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Service struct {
	products []ProductResponse
}

func (service Service) load() {
	jsonFile, err := os.Open("static/products.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened products.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &service.products)
}
