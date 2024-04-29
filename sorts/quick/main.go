package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, -2, -4, 9, 10}
	fmt.Println("Unsorted array: ", arr)
	arr = quickSort(arr)
	fmt.Println("Sorted array: ", arr)
}

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
