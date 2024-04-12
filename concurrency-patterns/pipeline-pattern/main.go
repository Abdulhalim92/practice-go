package main

import "fmt"

// Шаблон конвейера структурирует серию этапов обработки, каждый из
// которых выполняется одновременно. Данные проходят через эти этапы
// последовательно, что обеспечивает эффективное преобразование и
// обработку данных.
func main() {
	// Создаем начальный канал с некоторыми данными
	data := []int{1, 2, 3, 4, 5}
	input := make(chan int, len(data))

	for _, d := range data {
		input <- d
	}
	close(input)

	// Первый этап конвейера: удваивает входные значения
	doubleOutput := make(chan int)

	go func() {
		defer close(doubleOutput)

		for num := range input {
			doubleOutput <- num * 2
		}
	}()

	// Второй этап конвейера: возводит в квадрат удвоенные значения
	squareOutput := make(chan int)
	go func() {
		defer close(squareOutput)

		for num := range doubleOutput {
			squareOutput <- num * num
		}
	}()

	// Третий этап конвейера: печатает квадраты значений
	for result := range squareOutput {
		fmt.Println(result)
	}
}
