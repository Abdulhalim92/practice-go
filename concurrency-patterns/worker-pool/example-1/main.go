package main

import (
	"fmt"
	"sync"
)

// Задача: Конкурентное обработки изображений.
//
// Описание: У вас есть большое количество изображений, которые необходимо
// обработать. Обработка каждого изображения занимает определенное
// количество времени, и вы хотите конкурентно обрабатывать этот процесс, чтобы
// ускорить его выполнение. Для этого нужно использовать паттерн
// "worker pool".

// Размер worker pool.
const numWorkers = 5

// Исходные изображения для обработки.
var images = []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg", "image5.jpg"}

// Результаты обработки изображений.
var results = make(map[string]string)

// Реализация обработчика изображений.
func processImage(image string) string {
	// Здесь должна быть реализация обработки изображения.
	// Например:
	// processedImage := imageProcessingLogic(image)
	// return processedImage
	return "processed_" + image
}

// Рабочий, который будет обрабатывать изображения.
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for image := range jobs {
		result := processImage(image)
		results <- result
	}
}

func main() {
	// Создаем каналы для передачи работ и результатов.
	jobs := make(chan string, len(images))
	resultsChan := make(chan string, len(images))

	// Создаем wait group для ожидания завершения всех worker'ов.
	var wg sync.WaitGroup

	// Создаем worker.
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, resultsChan, &wg)
	}

	// Передаем изображения в канал jobs для обработки.
	for _, image := range images {
		jobs <- image
	}
	close(jobs)

	// Ожидаем завершения всех worker.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Получаем результаты обработки изображений.
	for result := range resultsChan {
		// Здесь можно сохранить результаты обработки.
		// Например:
		// results[imageName] = result
		fmt.Println("Processed image:", result)
	}
}
