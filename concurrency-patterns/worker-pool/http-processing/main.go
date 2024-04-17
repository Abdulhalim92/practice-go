package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Создайте HTTP-сервер, который принимает входящие HTTP запросы и
// добавляет их в очередь задач. Рабочие горутины могут извлекать
// запросы из очереди и обрабатывать их конкурентно.

// Количество рабочих программ (горутин)
const numWorkers = 3

// Рабочая функция для обработки запросов
func worker(id int, jobs <-chan *http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Здесь вы можете обрабатывать запрос
		// В этом примере мы просто выводим URL запроса
		fmt.Printf("Worker %d processing request for %s\n", id, job.URL.Path)
	}
}

func main() {
	// Создаем канал для задач
	jobs := make(chan *http.Request)

	// Создаем WaitGroup для отслеживания завершения всех рабочих программ
	var wg sync.WaitGroup

	// Создаем и запускаем рабочие программы
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Обработчик для всех входящих HTTP-запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Помещаем запрос в очередь задач
		jobs <- r
	})

	// Запускаем HTTP-сервер
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	// Ждем завершения всех рабочих программ
	wg.Wait()

	// Закрываем канал задач после завершения работы всех рабочих программ
	close(jobs)
}
