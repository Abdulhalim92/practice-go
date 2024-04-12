package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// Создайте HTTP-сервер, который принимает входящие HTTP запросы и
// добавляет их в очередь задач. Рабочие горутины могут извлекать
// запросы из очереди и обрабатывать их параллельно.

// Рабочая горутина
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %s\n", id, job)
		time.Sleep(2 * time.Second) // Иммитация длительной обработки
		fmt.Printf("Worker %d finished job %s\n", id, job)
		results <- job + "processed by worker " + fmt.Sprintf("%d", id)
	}
}

func main() {
	// Создание очередей для задач и результатов
	const numWorkers = 3
	jobs := make(chan string, 100)
	results := make(chan string, 100)

	// Создание WaitGroup
	var wg sync.WaitGroup

	// Добавление рабочих горутин в WaitGroup
	wg.Add(numWorkers)

	// Запуск рабочих горутин
	for i := 0; i < numWorkers; i++ {
		go worker(i+1, jobs, results, &wg)
	}

	// Обработчик HTTP запросов
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Чтение тела запроса
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Помешение запроса в очередь задач
		jobs <- string(body)
		fmt.Fprintln(w, "Job submitted successfully")
	})

	// Обработка результатов
	go func() {
		for result := range results {
			fmt.Println("Result:", result)
		}
	}()

	// Закрытие канала задач и ожидание завершения всех рабочих горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Запуск HTTP - сервера
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
