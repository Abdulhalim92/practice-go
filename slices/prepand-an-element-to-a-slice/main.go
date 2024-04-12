package main

import "fmt"

func main() {
	b := []int{2, 3, 4}
	b = append([]int{1}, b...)

	fmt.Println(b)
	// [1 2 3 4]
}
