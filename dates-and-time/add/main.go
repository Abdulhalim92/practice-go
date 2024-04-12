package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)

	date = date.AddDate(12, 1, 3)

	date = date.Add(time.Nanosecond * 1)
	date = date.Add(time.Microsecond * 5)
	date = date.Add(time.Millisecond * 5)
	date = date.Add(time.Second * 5)
	date = date.Add(time.Minute * 5)
	date = date.Add(time.Hour * 5)

	fmt.Println(date)
}
