package main

import (
	"fmt"
	"os"
)

func main() {
	// открыть файл
	f, err := os.Open("my-file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// прочить информацию о файле
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(info.Mode())

	fmt.Printf("Mode Numeric: %o\n", info.Mode())
	fmt.Printf("Mode Symbolic: %s\n", info.Mode())

	// изменение режима файла
	err = f.Chmod(0777)
	if err != nil {
		fmt.Println(err)
	}
}
