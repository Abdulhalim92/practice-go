package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	dbPortRaw := os.Getenv("DB_PORT")
	dbPort, err := strconv.ParseUint(dbPortRaw, 10, 16)
	if err != nil {
		log.Panicf("Impossible to parse database port number %s. Please double check the env variable DATABASE_PORT", dbPortRaw)
		return
	}

	fmt.Println(dbPort)
}
