package main

import "fmt"

// func delete(m map[Type]Type1, key Type)
func main() {

	// map literal
	var colors map[string]string
	colors = map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	fmt.Println(colors)
	// map[blue:#0000ff green:#00ff00 red:#ff0000]

	delete(colors, "red")
	fmt.Println(colors)
	// map[blue:#0000ff green:#00ff00]
}
