package main

import "fmt"

func main() {
	var x, z uint8
	x = 200                 // 11001000
	fmt.Printf("%08b\n", x) // Output: 11001000

	z = ^x

	// print in binary
	fmt.Printf("%08b\n", z) // Output: 00110111

	// print in base 10
	fmt.Printf("%d\n", z) // Output: 55
}
