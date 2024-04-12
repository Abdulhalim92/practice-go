package main

import (
	"fmt"
	"sync"
)

// Напишите функцию, которая принимает функцию и возвращает новую функцию.
// Эта новая функция вызывает переданную функцию только в случае, если
// результат еще не был сохранен, иначе возвращает сохраненный результат.

type MemoizedFunc func(int) int

func Memoize(f MemoizedFunc) MemoizedFunc {
	cache := make(map[int]int)
	var mutex sync.Mutex

	return func(n int) int {
		mutex.Lock()
		defer mutex.Unlock()

		if result, found := cache[n]; found {
			fmt.Println("From Cache")
			return result
		}

		result := f(n)
		cache[n] = result
		fmt.Println("Calculated")
		return result
	}
}

func ExampleFunction(x int) int {
	return x * x
}

func main() {
	// Срздаем мемоизированную версию функции
	memoizedExampleFunction := Memoize(ExampleFunction)

	fmt.Println(memoizedExampleFunction(5)) // Calculated
	fmt.Println(memoizedExampleFunction(5)) // From Cache
	fmt.Println(memoizedExampleFunction(6)) // Calculated
	fmt.Println(memoizedExampleFunction(6)) // From Cache
}
