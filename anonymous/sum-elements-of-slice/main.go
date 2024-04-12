package main

import "fmt"

// Напишите функцию, которая принимает слайс чисел и использует
// анонимную функцию для подсчета и возврата их суммы.
func countAndSum(slice []int) (int, int) {
	var (
		count = 0
		sum   = 0
	)

	func() {
		count = len(slice)

		for _, elem := range slice {
			sum += elem
		}
	}()

	return count, sum
}

func main() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	count, sum := countAndSum(mySlice)

	fmt.Printf("count elements of slice: %d and sum: %d", count, sum)
}
