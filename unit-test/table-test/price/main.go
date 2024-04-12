package main

import "fmt"

func TotalPrice(nights, rate, cityTax uint) uint {
	return nights * (rate + cityTax)
}

func main() {
	price := TotalPrice(3, 10000, 132)
	if price == 30396 {
		fmt.Println("function works")
	} else {
		fmt.Println("function is buggy")
	}
}
