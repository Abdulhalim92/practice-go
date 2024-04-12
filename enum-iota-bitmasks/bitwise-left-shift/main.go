package main

import "fmt"

func main() {
	var x, n, z uint8
	x = 200                 // 11001000
	fmt.Printf("%08b\n", x) // Output: 11001000

	// number of positions
	n = 1
	z = x << n
	// print in binary
	fmt.Printf("%08b\n", z) // Output: 10010000

	// print in base 10
	fmt.Printf("%d\n", z) // Output: 144

	fmt.Println("----------------------------------")
	var a, m, b uint16
	a = 200                 // 11001000
	fmt.Printf("%08b\n", a) // Output: 11001000

	// number of positions
	m = 1
	b = a << m
	// print in binary
	fmt.Printf("%08b\n", b) // Output: 10010000

	// print in base 10
	fmt.Printf("%d\n", b) // Output: 400
}
