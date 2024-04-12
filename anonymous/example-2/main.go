package main

import (
	"fmt"
	"reflect"
)

func main() {
	// func literal not executed
	myFunc := func() int {
		fmt.Println("I am a func literal")
		return 42
	}
	// Output: nothing
	// the function is not executed

	fmt.Println(reflect.TypeOf(myFunc()))
	// Output: func() int

	// func literal invoked
	funcValue := func() int {
		fmt.Println("I am a func literal invoked")
		return 42
	}()
	// Output: I am a func literal invoked
	// the function is executed
	fmt.Println(reflect.TypeOf(funcValue))
	// Output: int
}
