package main

import "fmt"

// s = append(s, 0)
// copy(s[i+1:], s[i:])
// s[i] = x
func main() {
	// objective put "C" at index 2
	// index:            0    1    2    3    4
	letters := []string{"A", "B", "D", "E", "F"}

	// 1) add an element to the end of the slice
	letters = append(letters, "")

	fmt.Println("after adding an element:", letters)

	// 2) copy letters[i:] to letters[i+1:]
	fmt.Println("before copying: dest -", letters[3:], "and src -", letters[2:])

	copy(letters[3:], letters[2:])

	fmt.Println("after copying:", letters)

	// 3) set "C" at index 2
	letters[2] = "C"

	fmt.Println(letters)

	second()
}

func second() {
	names := []string{"john", "jeane", "jean", "josh"}

	names = append(names, "")
	copy(names[2:], names[1:])
	names[1] = "joe"

	fmt.Println("second example: names:", names)
}
