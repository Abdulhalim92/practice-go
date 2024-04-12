package main

import "fmt"

// Напишите функцию, которая принимает строку и возвращает новую строку,
// в которой удалены все повторяющиеся символы
func main() {
	str := "hello world"
	fmt.Println("String with duplicates removed: ", removeDuplicates(str))
}

func removeDuplicates(s string) string {
	seen := make(map[rune]bool)
	var result []rune

	for _, char := range s {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}

	return string(result)
}
