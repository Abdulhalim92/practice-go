package main

import (
	"fmt"
	"strings"
)

// Напишите функцию, которая принимает строку и возвращает новую строку,
// в которой порядок слов обращен
func reverseWords(s string) string {
	words := strings.Fields(s)
	reversed := make([]string, len(words))

	for i, word := range words {
		reversed[len(words)-i-1] = word
	}

	return strings.Join(reversed, " ")
}

func main() {

	str := "Hello, World!"
	fmt.Println(reverseWords(str))
}
