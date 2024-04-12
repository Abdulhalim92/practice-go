package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	users := make([]User, 2)

	file, err := os.Open("./json-xml/problem/problem.json")
	defer func(file *os.File) { _ = file.Close() }(file)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		panic(err)
	}

	fmt.Println(users)
}
