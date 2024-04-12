package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ID   uint64   `json:"id"`
	Name string   `json:"name"`
	SKU  string   `json:"sku"`
	Cat  Category `json:"cat"`
}

type Category struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func main() {
	p := Product{
		ID:   42,
		Name: "Tea Pot",
		SKU:  "TP12",
		Cat: Category{
			ID:   2,
			Name: "Tea",
		},
	}

	bytes, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	bytesIndent, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytesIndent))

}
