package main

import "fmt"

func main() {
	// Поиск элемента в срезе не имеет встроенных функций поиска
	names := []string{"John", "Bob", "Claire", "Nik"}
	for i, name := range names {
		if name == "Claire" {
			fmt.Println("Claire found at index", i)
		}
	}
	// Output: Claire found at index 2
}
