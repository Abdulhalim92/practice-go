package main

import "fmt"

func main() {
	// Raw string
	raw := `spring rain:
browsing under an umbrella
at the picture-book store`
	fmt.Println(raw)

	// Interpreted string
	interpreted := "I love you"
	fmt.Println(interpreted)

	symbol := "µ"
	fmt.Println("\\xc2\\xb5")  // µ
	fmt.Println("\\u00b5")     // µ
	fmt.Println("\\U000000B5") // µ
	fmt.Printf("%x", symbol)   // c2b5
}
