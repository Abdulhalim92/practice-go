package main

import "fmt"

func main() {
	var x, y, z uint8
	x = 1 // 00000001
	y = 2 // 00000010
	z = x & y

	// print in binary
	fmt.Printf("%08b\n", z) // Output: 00000000

	// print in base 10
	fmt.Printf("%d\n", z) // Output: 0
}
