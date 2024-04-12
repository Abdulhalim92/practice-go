package main

import (
	"fmt"
	"os"
)

func main() {
	// Создание файла
	f, err := os.Create("./data-storage/file-save/simple/test.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// Открытие файла, если не существует, создавать
	f2, err := os.OpenFile("./data-storage/file-save/simple/test.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
}

// ls -al   --- визуализировать права доступа к файлу
