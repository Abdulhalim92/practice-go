package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Создайте HTTP-сервер, который принимает входящие HTTP запросы и
// добавляет их в очередь задач. Рабочие горутины могут извлекать
// запросы из очереди и обрабатывать их конкурентно.

type resultWithError struct {
	result result
	err    error
}

type result struct {
}

func worker(ctx context.Context, wg *sync.WaitGroup, jobs <-chan string, results chan<- resultWithError) {
	defer wg.Done()
	for j := range jobs {
		fmt.Printf("Enqueued job: %s\n", j)
		// Обработка задачи и добавление результата в канал результатов
		results <- resultWithError{
			result: result{},
			err:    nil,
		}
	}
}

func main() {
	// Создание контекста для управления рабочими горутинами
	ctx, cancel := context.WithTimeout(context.Background(), 0)
	defer cancel()

	// Создание каналов для передачи задач и результатов
	jobs := make(chan string, 100)
	results := make(chan resultWithError, 100)

	// Определение количества рабочих горутин
	numWorkers := 5

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Запуск рабочих горутин
	for i := 0; i < numWorkers; i++ {
		go worker(ctx, &wg, jobs, results)
	}

	// Обработчик HTTP запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			fmt.Printf("Method: %s\n", r.Method)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Декодирование JSON из тела запроса
		var reqData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
			fmt.Printf("Error: %v\n", err)
			http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
			return
		}

		log.Printf("Request: %v\n", reqData)

		// Добавление задачи в очередь
		jobs <- reqData["job"].(string)

		// Отправка успешного ответа
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "OK")
	})

	// Запуск HTTP-сервера в горутине
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("HTTP server error: %v\n", err)
		}
	}()

	// Ожидание завершения работы всех рабочих горутин
	go func() {
		wg.Wait()
		close(jobs)
		close(results)
	}()

	// Обработкк результатов
	for res := range results {
		if res.err != nil {
			fmt.Printf("Error: %v\n", res.err)
			continue
		}
		fmt.Printf("Result: %v\n", res.result)
	}
}
