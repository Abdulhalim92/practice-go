package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	err := os.WriteFile("/tmp/this-is-a-test", []byte("TEST"), 0777)
	if err != nil {
		panic(err)
	}

	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Printf("process took %s", elapsed)
}
