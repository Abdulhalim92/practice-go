package main

import "fmt"

func printer() func() {
	k := 1
	return func() {
		fmt.Printf("Print n. %d\n", k)
		k++
	}
}
func main() {
	p := printer() // func()
	p()
	// Print n. 1
	p()
	// Print n. 2
	p()
	// Print n. 3
}
