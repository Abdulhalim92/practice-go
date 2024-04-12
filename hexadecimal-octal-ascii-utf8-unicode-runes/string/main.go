package main

import "fmt"

// rune is an alias for int32
func main() {
	rusText := "тектс"
	fmt.Println(rusText[0], rusText[1]) // 209 130

	str1 := "Бухарь"
	str2 := "Буква"
	fmt.Println(str1 > str2) // true

	str1 = "Кот"
	str2 = "кот"
	fmt.Println(str1 == str2) // false

	str1 = "код"
	str2 = "кот"
	fmt.Println(str1 > str2) // false

	fmt.Println([]byte(str1)) // [208 186 208 190 208 180]
	fmt.Println([]byte(str2)) // [208 186 208 190 209 130]
}
