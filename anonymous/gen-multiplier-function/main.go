package main

import (
	"fmt"
	"reflect"
)

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	/*
		var double func(int)int
	*/
	triple := multiplier(3)
	/*
		var triple func(int)int
	*/
	fmt.Println(reflect.TypeOf(double)) // func(int)int
	fmt.Println(reflect.TypeOf(triple)) // func(int)int

	fmt.Println(double(5)) // Output: 10
	fmt.Println(triple(5)) // Output: 15
}
