package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFoo(t *testing.T) {
	expected := "Foo"
	actual := Foo()
	if expected != actual {
		t.Errorf("Expected %s do not match actual %s", expected, actual)
	}
}

func TestFoo2(t *testing.T) {
	assert.Equal(t, "Foo", Foo(), "they should be equal")
}
