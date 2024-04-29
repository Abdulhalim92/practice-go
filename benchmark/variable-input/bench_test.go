package variable_input

import (
	"fmt"
	"practice_go/benchmark/basic"
	"testing"
)

var result string

func BenchmarkConcatenation(b *testing.B) {
	var s string

	lengths := []int{2, 16, 128, 1024, 8192, 65536, 524288, 4194304, 16777216, 134217728}

	for _, l := range lengths {
		first := generateRandomString(l)
		second := generateRandomString(l)

		b.Run(fmt.Sprintf("ConcatenationJoin-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s = basic.ConcatenateJoin(first, second)
			}
			result = s
		})

		b.Run(fmt.Sprintf("ConcatenationBuffer-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				s = basic.ConcatenateBuffer(first, second)
			}
			result = s
		})
	}
}

// go test -bench=. -benchmem &>> benchmarkConcatenation.log
