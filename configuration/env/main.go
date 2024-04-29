package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	myVar := os.Getenv("MYVAR")
	myVar2 := os.Getenv("MYVAR2")
	fmt.Printf("myVar: %s\n", myVar)
	fmt.Printf("myVar2: %s\n", myVar2)

	// version 1 : env
	port, found := os.LookupEnv("DB_PORT")
	if !found {
		log.Fatal("impossible to start up, DB_PORT env var is mandatory")
	}
	portParsed, err := strconv.ParseUint(port, 10, 8)
	if err != nil {
		log.Fatalf("impossible to parse db port: %s", err)
	}
	log.Println(portParsed)

	// получение всех переменных среды
	fmt.Println(os.Environ())

	// устновление переменной среды
	err = os.Setenv("MYVAR3", "test3")
	if err != nil {
		panic(err)
	}
}

// export MYVAR=test && export MYVAR2=test2 && go run ./configuration/env/main.go
// export MYVAR=test && export MYVAR2=test2 && DB_PORT=5432 go run ./configuration/env/main.go
// Output: 2024/04/20 10:15:23 impossible to parse db port: strconv.ParseUint: parsing "5432": value out of range
// export MYVAR=test && export MYVAR2=test2 && DB_PORT=5 go run ./configuration/env/main.go
// Output: 2024/04/20 10:15:14 5
