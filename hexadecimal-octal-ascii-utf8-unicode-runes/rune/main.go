package main

import "fmt"

func main() {
	var aRune rune = 'Z'
	fmt.Printf("Unicode Code point of &#39;%c&#39;: %U\n",
		aRune,
		aRune,
	)

	var aByte byte = 'Z'
	fmt.Printf("Byte value of &#39;%c&#39;: %d\n",
		aByte,
		aByte,
	)
}
