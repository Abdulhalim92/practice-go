package main

import "fmt"

func main() {
	// first method: a[:0]
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a = a[:0]
	fmt.Println(a)
	// []
	fmt.Println(len(a))
	// 0
	fmt.Println(cap(a))
	// 10

	// second method: set the slice to nil
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b = nil
	fmt.Println(b)
	// []
	fmt.Println(len(b))
	// 0
	fmt.Println(cap(b))
	// 0
}
