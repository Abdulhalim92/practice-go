package main

import "fmt"

func main() {
	rusText := "текст"

	for index, char := range rusText {
		fmt.Printf("Руна: %#U с индексов %d\n", char, index)
	}
	fmt.Println("Длина строки:", len(rusText))

	engText := "text"
	for index, char := range engText {
		fmt.Printf("Руна: %#U с индексов %d\n", char, index)
	}
	fmt.Println("Длина строки:", len(engText))
}

//Руна: U+0442 'т' с индексом 0
//Руна: U+0435 'е' с индексом 2
//Руна: U+043A 'к' с индексом 4
//Руна: U+0441 'с' с индексом 6
//Руна: U+0442 'т' с индексом 8
//Длина строки: 10

//Руна: U+0074 't' с индексов 0
//Руна: U+0065 'e' с индексов 1
//Руна: U+0078 'x' с индексов 2
//Руна: U+0074 't' с индексов 3
//Длина строки: 4
