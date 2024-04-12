package tdd

import "testing"

// Разработка через тестирование (или TDD) — это метод разработки,
// при котором вы разрабатываете тесты до фактического написания
// программного обеспечения

func TestVowelCount(t *testing.T) {
	expected := uint(5)
	actual := VowelCount("I love you")
	if actual != expected {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}
