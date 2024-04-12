package main

import "fmt"

func main() {
	languages := []string{"Java", "PHP", "C"}
	fmt.Println("Capacity:", cap(languages))
	// Capacity: 3

	// cal function
	addGoFixed(&languages)

	fmt.Println("Capacity:", cap(languages))
	// Capacity: 6
	fmt.Println(languages)
	// [Java PHP C Go]
}

func addGoFixed(languages *[]string) {
	*languages = append(*languages, "Go")
	fmt.Println("in function, capacity:", cap(*languages))
}
