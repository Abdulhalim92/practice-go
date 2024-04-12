package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "текст"
	ln := len(str)                             // 10
	correctLen1 := utf8.RuneCountInString(str) // 5
	correctLen2 := len([]rune(str))            // 5

	fmt.Println(ln)
	fmt.Println(correctLen1)
	fmt.Println(correctLen2)
}
