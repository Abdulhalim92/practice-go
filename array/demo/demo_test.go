package demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// updating an array with a value
func TestUpdateArray1(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray1(testArray)
	assert.Equal(t, NewValue, testArray[0])
}

// updating an array with a pointer
func TestUpdateArray2(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	UpdateArray2(&testArray)
	assert.Equal(t, NewValue, testArray[0])
}

// getting a copy of an array
func TestArrayCopy(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	newCopy := testArray
	testArray[1] = "updated"
	assert.Equal(t, "updated", newCopy[1])
}

// getting the address of an array
func TestArrayReference(t *testing.T) {
	testArray := [2]string{"Value1", "Value2"}
	reference := &testArray
	testArray[1] = "updated"
	assert.Equal(t, "updated", reference[1])
}
