package main

import (
	"errors"
	"fmt"
)

var errSentinel = errors.New("test")

func main() {
	err := foo()
	if errors.Is(err, errSentinel) {
		fmt.Println(err)
		fmt.Println(errSentinel)
		fmt.Println("errSentinel detected in the err chain with errors.Is")
	}
	if err == errSentinel {
		fmt.Println(err)
		fmt.Println(errSentinel)
		fmt.Println("errSentinel detected in the error chain by ==")
	}
}

func foo() error {
	return fmt.Errorf("error: %w", bar())
}

func bar() error {
	return errSentinel
}
