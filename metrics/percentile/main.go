package main

import "fmt"

func nthPercentile(sortedArray []int, n int) float64 {
	// Вычисляем индекс элемента, который представляет собой n-й перцентиль
	index := float64(n) / 100 * float64(len(sortedArray)-1)

	// Если индекс - целое число, то перцентиль - это значение в массиве в этом индексе
	if index == float64(int(index)) {
		return float64(sortedArray[int(index)])
	}

	// Если индекс - дробное число, то перцентиль - это линейная интерполяция между ближайшими значениями в массиве
	lowerIndex := int(index)
	upperIndex := lowerIndex + 1
	lowerValue := float64(sortedArray[lowerIndex])
	upperValue := float64(sortedArray[upperIndex])
	percentile := lowerValue + (upperValue-lowerValue)*(index-float64(lowerIndex))

	return percentile
}

func main() {
	// Пример сортированного массива
	sortedArray := []int{200, 200, 300, 300, 1000}

	// Вычисляем 50-й перцентиль (медиана)
	nth := 99
	result := nthPercentile(sortedArray, nth)
	fmt.Printf("%d-й перцентиль: %.2f\n", nth, result)
}
