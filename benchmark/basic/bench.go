package basic

import (
	"bytes"
	"strings"
)

func ConcatenateBuffer(first string, second string) string {
	var buffer bytes.Buffer
	buffer.WriteString(first)
	buffer.WriteString(second)
	return buffer.String()
}

func ConcatenateJoin(fisrt string, second string) string {
	return strings.Join([]string{fisrt, second}, "")
}
