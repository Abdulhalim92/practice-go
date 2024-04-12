package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Создайте HTTP-сервер, который принимает входящие HTTP запросы и
// добавляет их в очередь задач. Рабочие горутины могут извлекать
// запросы из очереди и обрабатывать их параллельно.

// TaskQueue представляет собой очередь задач
type TaskQueue struct {
	queue chan string
}

// NewTaskQueue возвращает новый экземпляр TaskQueue с заданным размером очереди
func NewTaskQueue(size int) *TaskQueue {
	return &TaskQueue{
		queue: make(chan string, size),
	}
}

// Add добавляет задачу в очередь
func (q *TaskQueue) Add(task string) {
	q.queue <- task
}

// Process начинает обработку задач из очереди
func (q *TaskQueue) Process() {
	for {
		select {
		case task := <-q.queue:
			// Здесь может быть ваш код обработки задачи
			fmt.Println("Processing task:", task)
			time.Sleep(2 * time.Second) // Имитация длительной обработки
			fmt.Println("Task", task, "completed")
		}
	}
}

func main() {
	// Создаем очередь задач
	tq := NewTaskQueue(10)

	// Запускаем горутину для обработки задач из очереди
	go tq.Process()

	// Обработчик HTTP запросов
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Читаем тело запроса
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Преобразуем байты тела запроса в строку и добавляем в очередь задач
		tq.Add(string(body))

		// Возвращаем ответ клиенту
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Task added to queue"))
	})

	// Запускаем HTTP-сервер на порту 8080
	fmt.Println("Starting server on port 8080")
	log.Fatalf("http.ListenAndServe: %v", http.ListenAndServe(":8080", nil))
}
