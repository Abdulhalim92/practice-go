package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	multiply(s, 2)
	fmt.Println(s)
	// [2,4,6]
}

func multiply(slice []int, factor int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = slice[2] * factor
	}
}
