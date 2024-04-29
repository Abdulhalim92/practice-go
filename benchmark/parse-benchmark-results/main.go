package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	b, err := os.ReadFile("./benchmark/variable-input/benchmarkConcatenation.log")
	if err != nil {
		panic(err)
	}

	benchmarkResult := string(b)

	fmt.Println(benchmarkResult)

	regexBench := regexp.MustCompile(`([a-zA-Z]*)-(\d+)-.* (\d+\.?\d+?)[\t]ns.*[\t](\d+)[\t]B.* (\d+) allocs`)
	matches := regexBench.FindAllStringSubmatch(benchmarkResult, -1)

	fmt.Println("benchmarkedFunction,stringLen,nsPerOp,bytesPerOp,mallocs")

	for _, m := range matches {
		fmt.Printf("%s,%s,%s,%s,%s\n", m[1], m[2], m[3], m[4], m[5])
	}
}
