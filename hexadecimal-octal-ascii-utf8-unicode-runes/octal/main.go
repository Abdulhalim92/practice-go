package main

import "fmt"

func main() {
	n2 := 0x9F4
	fmt.Printf("Decimal : %d\n", n2)

	// n3 is represented using the octal numeral system
	n3 := 02454
	// alternative : n3 := 0o2454

	// convert in decimal
	fmt.Printf("decimal: %d\n", n3)

	// n4 is represented using the decimal numeral system
	n4 := 1324
	// output n4 (decimal) in octal
	fmt.Printf("octal: %o\n", n4)
	// output n4 (decimal) in octal (with a 0o prefix)
	fmt.Printf("octal with prefix : %O\n", n4)
}
