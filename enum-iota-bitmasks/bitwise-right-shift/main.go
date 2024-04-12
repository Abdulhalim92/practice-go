package main

import "fmt"

func main() {
	var x, n, z uint8
	x = 200                 // 11001000
	fmt.Printf("%08b\n", x) // Output: 11001000

	// number of positions
	n = 3
	z = x >> n
	// print in binary
	fmt.Printf("%08b\n", z) // Output: 0011001

	// print in base 10
	fmt.Printf("%d\n", z) // Output: 25
}
