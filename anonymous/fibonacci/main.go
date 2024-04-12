package main

import "fmt"

// Напишите функцию, которая возвращает анонимную функцию,
// генерирующую следующий элемент последовательности Фибоначчи
// при каждом вызове.

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		f2 := f1 + f0
		f0 = f1
		f1 = f2
		return f0
	}
}

func main() {
	f := fibonacci()
	fmt.Println(f()) // Output: 1
	fmt.Println(f()) // Output: 1
	fmt.Println(f()) // Output: 2
	fmt.Println(f()) // Output: 3
	fmt.Println(f()) // Output: 5
}
