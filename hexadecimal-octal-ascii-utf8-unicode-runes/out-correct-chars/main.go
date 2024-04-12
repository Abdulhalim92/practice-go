package main

import "fmt"

func main() {
	str := "текст"
	for _, sym := range str {
		fmt.Printf("%c ", sym)
	}
}
