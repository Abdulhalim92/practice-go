package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, -2, -4, 9, 10}
	fmt.Println(simpleSort(arr))
}

func simpleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
