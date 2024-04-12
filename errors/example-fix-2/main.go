package main

import (
	"fmt"
	"log"
	"practice_go/errors/example-fix-2/config"
)

func main() {
	confData, err := config.Load()
	if err != nil {
		log.Fatalf("Impossible to load application config because: %s", err)
	}
	fmt.Println(confData)
}
