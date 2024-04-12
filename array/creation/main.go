package main

import "fmt"

func main() {
	var myArray [2]int
	myArray[0] = 156
	myArray[1] = 147
	fmt.Println(myArray)

	// краткий способ:
	myArray2 := [2]int{156, 147}
	fmt.Println(myArray2)

	// вычисление длины компилятором:
	c := [...]float64{3.14, 2.71, 1.62}
	fmt.Println(c)

	// не объявлены значения:
	d := [3]int{}
	fmt.Println(d)

	fmt.Println(len(d))
	fmt.Println(cap(d))
}
