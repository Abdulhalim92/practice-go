package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type ReadingError struct {
	Filename string
	IOError  error
}

func (e *ReadingError) Error() string {
	return fmt.Sprintf("impossible to read file")
}

func (e *ReadingError) Unwrap() error {
	return e.IOError
}

type WritingError struct {
	Filename string
	IOError  error
}

func (e *WritingError) Error() string {
	return fmt.Sprintf("impossible to write file")
}

func (e *WritingError) Unwrap() error {
	return e.IOError
}

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

// func As(err error, target interface{}) bool
// Эту функцию errors.As можно использовать для определения того,
// возвращается ли в цепочке ошибок ошибка типа X
func main() {
	err := transferFileContacts("test.csv")

	var readingErr *ReadingError
	if errors.As(err, &readingErr) {
		log.Fatalf("error of reading occured: %s: %s", readingErr, readingErr.Unwrap())
	}

	var writingErr *WritingError
	if errors.As(err, &writingErr) {
		log.Fatalf("error of writing occured: %s: %s", writingErr, writingErr.Unwrap())
	}

	log.Println("transfer done")
}
