package main

import (
	"fmt"
	"reflect"
)

type employee struct {
	name     string
	genre    string
	position string
}

func main() {
	employees := make(map[string]employee)

	employees["first"] = employee{
		name:     "John Smith",
		genre:    "M",
		position: "CEO",
	}

	walter := employees["first"]
	fmt.Println(walter)
	// {John Smith M CEO}

	ghost := employees["second"]
	fmt.Println(ghost)
	// {   }

	fmt.Println(reflect.TypeOf(ghost))
	// bitwise-or.employee

	foo, ok := employees["third"]
	if !ok {
		fmt.Println("failed to find the employee")
	} else {
		fmt.Println(foo)
	}
	// failed to find the employee
}
