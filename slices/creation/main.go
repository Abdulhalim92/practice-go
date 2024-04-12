package main

import "fmt"

func main() {
	// создание слайса без инициализации
	s := make([]int, 3)
	s[0] = 12
	s[2] = 3
	fmt.Println(s)

	// создание слайса с инициализацией
	s2 := []int{3, 4, 5}
	fmt.Println(s2)

	// создание пустого слайса
	var s3 []int
	fmt.Println(s3)

}
