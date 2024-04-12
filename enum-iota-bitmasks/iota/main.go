package main

import "fmt"

type HTTPMethod int

const (
	GET     HTTPMethod = iota
	POST    HTTPMethod = iota
	PUT     HTTPMethod = iota
	DELETE  HTTPMethod = iota
	PATCH   HTTPMethod = iota
	HEAD    HTTPMethod = iota
	OPTIONS HTTPMethod = iota
	TRACE   HTTPMethod = iota
	CONNECT HTTPMethod = iota
)

type TestEnum int

const (
	First TestEnum = iota + 4
	Second
)

func main() {
	fmt.Println(PUT)    // Output: 2
	fmt.Println(Second) // Output: 5
}
