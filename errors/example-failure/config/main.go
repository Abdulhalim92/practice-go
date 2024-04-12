package main

import (
	"fmt"
	"os"
)

func load() []byte {
	data, err := os.ReadFile("./errors/example-failure/config/conf.json")
	fmt.Println(err)
	// open ./errors/example-failure/config/conf.json: no such file or directory
	return data
}

func Print() {
	fmt.Println(string(load()))
}

func main() {
	Print()
}
