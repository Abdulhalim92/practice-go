package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Узнать текущее время
	now := time.Now()
	fmt.Printf("Current time: %02d:%02d:%02d\n", now.Hour(), now.Minute(), now.Second())
	// Узнать текущее время (полный)
	fmt.Printf("%s\n", now)

	// Узнать текущий миллисекундный UNIX-время
	log.Println(now.Unix())

}
