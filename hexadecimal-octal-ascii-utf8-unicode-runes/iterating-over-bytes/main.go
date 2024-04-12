package main

import "fmt"

func main() {
	rusText := "текст"
	for i := 0; i < len(rusText); i++ {
		fmt.Printf("%v ", rusText[i])
	}
	fmt.Println()
	// цикл для перевода 10-ричных значений байтов в 16-ричные
	for i := 0; i < len(rusText); i++ {
		fmt.Printf("%x ", rusText[i])
	}
}
