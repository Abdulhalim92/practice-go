package main

import "fmt"

// синтаксис анонимных функций: func(){...}()
func main() {
	// an anonymous function
	func() {
		fmt.Println("I am an anonymous function")
	}()
}
