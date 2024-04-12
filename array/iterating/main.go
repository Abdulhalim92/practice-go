package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}

	// iteration over an array
	for index, element := range myArray {
		fmt.Printf("element at index %d is %d\n", index, element)
	}

	// iteration over an array without index
	for _, element := range myArray {
		fmt.Printf("element is %d\n", element)
	}

	// iteration over an array with index
	for index := range myArray {
		fmt.Printf("element at index %d is %d\n", index, myArray[index])
	}

	// iteration over an array with index
	for i := 0; i < len(myArray); i++ {
		fmt.Printf("element at index %d is %d\n", i, myArray[i])
	}

	// iteration over an array with index descending
	for i := len(myArray); i >= 0; i-- {
		fmt.Printf("element at index %d is %d\n", i, myArray[i])
	}
}
