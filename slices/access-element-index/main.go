package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40}
	fmt.Println(a[2])
	// Output: 30
	fmt.Println(a[8])
	// panic: runtime error: index out of range
}
