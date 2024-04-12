package main

import "sort"

// Search for an index of a specific value using binary search
// func Search(n int, f func(int) bool) int
//
// sort a slice by using the provided function (less)
// func Slice(slice interface{}, less func(i, j int) bool)
func main() {
	scores := []int{10, 89, 76, 3, 20, 12}

	// ascending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] < scores[j]
	})
	// Output: [3 10 12 20 76 89]

	// descending
	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})
	// Output: [89 76 20 12 10 3]
}
