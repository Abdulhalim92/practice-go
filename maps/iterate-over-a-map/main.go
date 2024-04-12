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

	for key, value := range colors {
		fmt.Printf("Key: %s - Value: %s\n", key, value)
	}
	//Key: red - Value: #ff0000
	//Key: green - Value: #00ff00
	//Key: blue - Value: #0000ff

}
