package main

import "fmt"

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

	// map constructor
	worldCupWinners := map[int]string{
		1930: "Uruguay",
		1934: "Italy",
		1938: "Italy",
		1950: "Uruguay",
	}
	fmt.Println(worldCupWinners)
	// map[1930:Uruguay 1934:Italy 1938:Italy 1950:Uruguay]

	// empty map
	emptyMap := make(map[string]string)
	fmt.Println(emptyMap)
	// map[]
}
