package main

import "fmt"

func main() {
	test := append([]int{10, 20}, []int{30, 40, 50}...)
	fmt.Println(test)
	// [10 20 30 40 50]
}
