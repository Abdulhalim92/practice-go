package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

func main() {
	// вывод с помощью функции maxInt64
	var a, b int64 = 42, 23
	fmt.Println(maxInt64(a, b))

	// вывод с помощью функции maxInt32
	var c, d int32 = 12, 376
	fmt.Println(maxInt32(c, d))

	// вывод с помощью функции maxInt8
	var e, f int8 = 12, 3
	fmt.Println(maxInt8(e, f))

	// вывод с помощью функции maxGeneric
	fmt.Println(maxGeneric[int64](12, 3))
	fmt.Println(maxGeneric(12, 3))

	// вывод с помощью функции maxEmptyInterface
	fmt.Println(maxEmptyInterface(12, 3))
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maxInt32(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func maxInt8(a, b int8) int8 {
	if a > b {
		return a
	}
	return b
}

func maxGeneric[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func maxEmptyInterface(a, b interface{}) interface{} {
	if a > b {
		return a
	}
	return b
}
