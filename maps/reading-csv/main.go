package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type employee struct {
	name     string
	genre    string
	position string
}

func main() {
	employees := make(map[string]employee)

	file, err := os.Open("./maps/reading-csv/users.csv")
	if err != nil {
		log.Fatalf("impossible to open file: %s", err)
	}
	defer func(file *os.File) { _ = file.Close() }(file)

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		employee := employee{
			name:     record[1],
			genre:    record[2],
			position: record[3],
		}
		employees[record[0]] = employee
	}
	fmt.Println(employees)
	// map[alice:employee{name:alice genre:male position:developer} bob:employee{name:bob genre:male position:developer} carol:employee{name:carol genre:female position:designer}]
}
