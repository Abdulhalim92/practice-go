package tdd

var vowels = map[string]bool{
	"A": true,
	"E": true,
	"I": true,
	"O": true,
	"U": true,
	"a": true,
	"e": true,
	"i": true,
	"o": true,
	"u": true,
}

func VowelCount(sentence string) uint {
	var count uint
	for _, char := range sentence {
		if vowels[string(char)] {
			count++
		}
	}
	return count
}
