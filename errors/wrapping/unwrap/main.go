package main

import "fmt"

type ReadingError struct {
	IOError  error
	Filename string
}

// implement the error interface
func (e *ReadingError) Error() string {
	return fmt.Sprintf("an error occurred while attempting to read the file %s", e.Filename)
}

// Unwrap will implement the Unwrap method
func (e *ReadingError) Unwrap() error {
	return e.IOError
}
