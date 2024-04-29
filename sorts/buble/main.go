package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, -2, -4, 9, 10}
	fmt.Println(bubbleSort(arr))
}

func bubbleSort(arr []int) []int {
	var tmp, k int
	n := len(arr)

	for n > 1 {
		k = 0
		for i := 1; i < n; i++ {
			if arr[i] < arr[i-1] {
				tmp = arr[i-1]
				arr[i-1] = arr[i]
				arr[i] = tmp
				k = i
			}
		}
		n = k
	}

	return arr
}
