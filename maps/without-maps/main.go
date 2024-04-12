package main

import "fmt"

type testScore struct {
	studentName string
	score       uint8
}

func main() {
	results := []testScore{
		{"John Doe", 20},
		{"Patrick Smith", 30},
		{"Jane Doe", 40},
		{"Mark Smith", 50},
		{"Bob Ferris", 60},
		{"Claire Novalingua", 70},
	}

	fmt.Println(results)

	// not optimal solution
	// сложность алгоритма O(n)
	for _, result := range results {
		if result.studentName == "Claire Novalingua" {
			fmt.Println("Score Found: ", result.score)
		}
	}
}
