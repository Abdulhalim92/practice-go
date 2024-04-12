package main

import (
	"fmt"
	"log"
	"os"
)

func transferFileContacts(filename string) error {
	contents, err := os.ReadFile(filename)
	if err != nil {
		// wrapping the error with fmt.Errorf
		return fmt.Errorf("during file transfer impossible to open source file: %w", err)
	}

	err = os.WriteFile("/tmp/file-contents.txt", contents, 0644)
	if err != nil {
		// wrapping the error with fmt.Errorf
		return fmt.Errorf("during file transfer impossible to write file: %w", err)
	}

	return nil
}

func main() {
	err := transferFileContacts("./errors/example-failure/config/conf.json")
	if err != nil {
		log.Printf("error occurred: %s", err)
	}
}
