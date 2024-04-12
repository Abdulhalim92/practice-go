package main

import (
	"fmt"
	"time"
)

func main() {
	timeParse := "2019-02-15T07:33-05:00"
	parse, err := time.Parse("2006-01-02T15:04-07:00", timeParse)
	if err != nil {
		panic(err)
	}

	fmt.Println(parse)
}
