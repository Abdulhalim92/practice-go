package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// func Is(err, target error) bool
// С помощью функции Is из пакета ошибок вы можете определить,
// произошла ли конкретная ошибка
func main() {
	file, err := os.Open("test.csv")
	defer func(file *os.File) { _ = file.Close() }(file)
	if err != nil {
		log.Printf("impossible to open file: %s", err)
		return
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		// use Is instead of an equality comparison
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record)
	}
}
