package main

import "fmt"

func main() {
	var x, y, z uint8
	x = 200                 // 11001000
	fmt.Printf("%08b\n", x) // Output: 11001000

	y = 100                 // 01100100
	fmt.Printf("%08b\n", y) // Output: 01100100

	z = x | y

	// print in binary
	fmt.Printf("%08b\n", z) // Output: 11101100

	// print in base 10
	fmt.Printf("%d\n", z) // Output: 236
}
