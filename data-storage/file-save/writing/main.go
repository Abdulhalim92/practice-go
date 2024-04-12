package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	s := [][]string{
		{"age", "genre", "name"},
		{"23", "M", "Hendrick"},
		{"65", "F", "Stephany"},
	}

	// Открыть или создать файл
	f, err := os.Create("my-file.csv")
	if err != nil {
		fmt.Println(err)
		return
	}

	var buffer bytes.Buffer

	// итерация по срезу
	for _, data := range s {
		buffer.WriteString(fmt.Sprintf("%s,%s,%s,\n", data[0], data[1], data[2]))
	}

	n, err := f.Write(buffer.Bytes())
	fmt.Printf("%d bytes written\n", n)
	if err != nil {
		fmt.Println(err)
		return
	}
}
