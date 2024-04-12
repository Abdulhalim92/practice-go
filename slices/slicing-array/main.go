package main

import "fmt"

func main() {
	// инициаилизация массива
	customers := [4]string{"John", "Paul", "George", "Ringo"}

	// срез массива
	customersSlice := customers[1:3]
	fmt.Println(customersSlice)

	// срез массива с до последней позиции
	customersSlice2 := customers[1:]
	fmt.Println(customersSlice2)

	// срез массива с начала до первой позиции
	customersSlice3 := customers[:]
	fmt.Println(customersSlice3)

	// модификация оригинального массива
	customers[0] = "Pete"
	fmt.Println("after modification of original array: ")
	fmt.Println(customersSlice3)
}
