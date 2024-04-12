package main

import "fmt"

// func append(slice []Type, elems ...Type) []Type
func main() {
	a := []int{10, 20, 30, 40}
	a = append(a, 50)
	fmt.Println(a)
	// [10, 20, 30, 40, 50]

	s := []uint{10, 20, 30, 40}
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 50)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 60)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 70)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 80)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	s = append(s, 90)
	fmt.Printf("Length : %d - Capacity : %d\n", len(s), cap(s))
	//Length : 4 - Capacity : 4
	//Length : 5 - Capacity : 8
	//Length : 6 - Capacity : 8
	//Length : 7 - Capacity : 8
	//Length : 8 - Capacity : 8
	//Length : 9 - Capacity : 16

	// потребляет больше производительности, из за того, что создается новый
	// базовый массив и данные копируются
	grow1()
	// потребляет меньше производительности, из за того, что не создается
	// новый базовый массив
	grow2()
}

func grow1() {
	s := []uint{10, 20, 30, 40}
	s = append(s, 50)
	s = append(s, 60)
	s = append(s, 70)
	s = append(s, 80)
	s = append(s, 90)
}

func grow2() {
	s := make([]uint, 9)
	s = append(s, 10, 20, 30, 40)
	s = append(s, 50)
	s = append(s, 60)
	s = append(s, 70)
	s = append(s, 80)
	s = append(s, 90)
}
