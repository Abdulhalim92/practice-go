package main

import "fmt"

// type byte is alias of uint8
func main() {
	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	fmt.Println(b)
}
