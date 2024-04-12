package main

import (
	"encoding/json"
	"fmt"
)

// Product omitempty - игнорирует пустые поля
// Знак "-"  пропустит поле с указанным тегом
type Product struct {
	ID          uint64 `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"-"`
}

func main() {
	p := Product{ID: 42, Description: "This is the product"}
	data, err := json.MarshalIndent(p, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
