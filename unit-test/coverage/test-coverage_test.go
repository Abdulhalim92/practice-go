package coverage

import "testing"

// go test -coverprofile profileCount -covermode count unit-test/coverage/test-coverage.go
// go tool cover -func=profileCount
// cat profileCount | grep total
// go test -coverprofile profileAtomic -covermode atomic unit-test/coverage/test-coverage.go
// go tool cover -func=profileAtomic
func TestBazBaz(t *testing.T) {
	expected := 3
	actual := BazBaz(3)
	if actual != expected {
		t.Errorf("Expected %d does not match actual %d", expected, actual)
	}
}
