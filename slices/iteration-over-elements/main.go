package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Bob", "Claire", "Nik"}
	for index, name := range names {
		fmt.Println("Element at index", index, "=", name)
	}

	// Warning : there is a trap!
	for _, name := range names {
		name = strings.ToUpper(name)
	}
	fmt.Println(names)
	// [John Bob Claire Nik]

	for i := range names {
		names[i] = strings.ToUpper(names[i])
	}
	fmt.Println(names)
	// [JOHN BOB CLAIRE NIK]
}
