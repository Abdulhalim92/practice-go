package main

import (
	"fmt"
	"strings"
)

func main() {
	// количество итераций может быть в разы больше
	sb := strings.Builder{}
	for i := 0; i < 8; i++ {
		sb.WriteString("q")
	}
	sb.WriteString("end")
	fmt.Println(sb.String())
}
