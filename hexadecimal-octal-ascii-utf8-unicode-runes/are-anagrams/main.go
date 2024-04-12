package main

import (
	"fmt"
	"sort"
	"strings"
)

// Напишите функцию, которая принимает две строки и возвращает true,
// если они являются анаграммами (имеют одинаковые символы в разном порядке),
// и false в противном случае.
func areAnagrams(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	s1Runes := []rune(s1)
	s2Runes := []rune(s2)

	sort.Slice(s1Runes, func(i, j int) bool {
		return s1Runes[i] < s1Runes[j]
	})
	sort.Slice(s2Runes, func(i, j int) bool {
		return s2Runes[i] < s2Runes[j]
	})

	return string(s1Runes) == string(s2Runes)
}

func main() {
	str1 := "listen"
	str2 := "silent"

	if areAnagrams(str1, str2) {
		fmt.Println("The strings are anagrams")
	} else {
		fmt.Println("The strings are not anagrams")
	}
}
