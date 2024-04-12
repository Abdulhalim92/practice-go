package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("%s\n", now)

	loc, err := time.LoadLocation("Asia/Dushanbe")
	if err != nil {
		panic(err)
	}

	nowDU := now.In(loc)
	fmt.Printf("%s\n", nowDU)
}
