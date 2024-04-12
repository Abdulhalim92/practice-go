package main

import "fmt"

func main() {
	var x uint8
	x = 1
	fmt.Printf("%08b\n", x) // Output: 00000001
	x = 2
	fmt.Printf("%08b\n", x) // Output: 00000010
	x = 255
	fmt.Printf("%08b\n", x) // Output: 11111111
}
