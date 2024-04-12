package main

import "fmt"

// func copy(dst, src []Type) int
func main() {
	// destination > source
	a := []int{10, 20, 30, 40}
	b := []int{1, 1, 1, 1, 1}
	copy(b, a)
	fmt.Printf("a: %v - b: %v\n", a, b)
	// Output: a:= [10, 20, 30, 40] - b: [10, 20, 30, 40, 1]

	// source > destination
	a = []int{10, 20, 30, 40}
	b = []int{1, 1}
	copy(b, a)
	fmt.Printf("a: %v - b: %v\n", a, b)
	// Output: a: [10, 20, 30, 40] - b: [10, 20]

	// source = destination
	a = []int{10, 20, 30, 40}
	b = make([]int, 4)
	copy(b, a)
	fmt.Printf("a: %v - b: %v\n", a, b)
	// Output: a: [10, 20, 30, 40] - b: [10, 20, 30, 40]
}
