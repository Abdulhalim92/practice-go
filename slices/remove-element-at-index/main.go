package main

import "fmt"

// a = append(a[:i], a[i+1:]...) - удаление i - го элемента среза
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	nums = append(nums[:8], nums[9:]...)
	fmt.Println(nums)
	// [1 2 3 4 5 6 7 8 10]
}
