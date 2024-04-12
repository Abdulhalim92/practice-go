package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	env := os.Getenv("MYENV")
	fmt.Println(env)
}
