package main

import (
	"fmt"
	"os"
	"testing"
)

// go test unit-test/table-test/price/arg_test.go -args baz
func TestArgs(t *testing.T) {
	arg1 := os.Args[1]
	if arg1 != "baz" {
		fmt.Println(arg1)
		t.Errorf("Expected %s does not match actual %s", "baz", arg1)
	}
}
