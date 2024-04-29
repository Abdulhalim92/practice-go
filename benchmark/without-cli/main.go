package main

import (
	"fmt"
	"practice_go/benchmark/basic"
	"testing"
)

func main() {
	res := testing.Benchmark(BenchmarkConcatenateBuffer)
	fmt.Printf("Memory allocations: %d\n", res.MemAllocs)
	fmt.Printf("Number of bytes allocated: %d\n", res.Bytes)
	fmt.Printf("Number of run: %d\n", res.N)
	fmt.Printf("Time taken: %s\n", res.T)
}

var result string

func BenchmarkConcatenateBuffer(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = basic.ConcatenateBuffer("test2", "test3")
	}

	result = s
}

// go test -bench Join -- запустит все функции тестирования, содержащие Join
// go test -bench=. -- запустит все функции тестирования
// go test -bench BenchmarkConcatenateJoin -benchtime 5s -- контролирует время тестирования
