package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type GenericMap[K constraints.Ordered, V constraints.Integer] map[K]V

func (m GenericMap[K, V]) sum() V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := GenericMap[string, int]{
		"foo": 42,
		"bar": 44,
	}
	fmt.Println(m.sum())

	m2 := GenericMap[float32, uint8]{
		12.5: 0,
		2.2:  23,
	}
	fmt.Println(m2.sum())
	/*
		Не будет работать:
		m3 := map[string]uint8{
			"foo": 10,
		}
		fmt.Println(m3.sum())
	*/

}
