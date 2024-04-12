package main

import "fmt"

func main() {
	languages := []string{"Java", "PHP", "C"}
	fmt.Println("Capacity:", cap(languages))
	// Capacity: 3

	// cal function
	addGo(languages)

	fmt.Println("Capacity:", cap(languages))
	// Capacity: 3
	fmt.Println(languages)
	// [Java PHP C]
}

func addGo(languages []string) {
	languages = append(languages, "Go")
	fmt.Println("in function, capacity:", cap(languages))
}
