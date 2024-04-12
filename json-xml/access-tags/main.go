package main

import (
	"fmt"
	"reflect"
)

type Product struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name,omitempty"`
	Description string `json:"-"`
}

func main() {
	p := Product{ID: 32}
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("field Name : %s\n", t.Field(i).Name)
		fmt.Printf("field Tag : %s\n", t.Field(i).Tag)
	}

	fmt.Println("-----------------------------------------")

	// Get() получение тега
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Tag.Get("json"))
	}

	fmt.Println("------------------------------------------")

	// Lookup() поиск тега
	for i := 0; i < t.NumField(); i++ {
		if tagValue, ok := t.Field(i).Tag.Lookup("test"); ok {
			fmt.Println(tagValue)
		} else {
			fmt.Println("no tag test")
		}
	}
}
