package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "привет, мир!"
	strings.Count(str, "и") // 2
	fmt.Println(str)

	// последний аргумент - количество замен
	// если он < 0, то лимита на количество замен нет
	str = "привет, мир!!!"
	a := strings.Replace(str, "!", "?", -1) // "привет, мир???"
	b := strings.Replace(a, "?", "!", 2)    // "привет, мир!!?"
	fmt.Println(a)
	fmt.Println(b)

	str = "привет, мир!!!"
	c := strings.Split(str, ",") // [привет  мир!!!]
	fmt.Println(c[0])            // привет

	str2 := []string{"01", "01", "2024"}
	strings.Join(str2, "-") // 01-01-2024
	fmt.Println(str2)

	str = "Просто Строка"
	strings.ToLower(str) // просто строка
	fmt.Println(str)
	strings.ToUpper(str) // ПРОСТО СТРОКА
	fmt.Println(str)
}

// Подсчет вхождений символа в строку – strings.Count(s, substr string) int
// Замена символов – strings.Replace(s, old, new string, n int) string
// Разбиение символов по разделителю sep – strings.Split(s, sep string) []string
// Конкатенация – strings.Join(s, sep string) string
// Преобразование символов к нижнему или верхнему регистру – strings.ToLower(s string)
// или strings.ToUpper(s string)
