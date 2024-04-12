package main

import (
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	const numRequests = 30

	// Создание каналов для задач и результатов
	jobs := make(chan string, numRequests)
	results := make(chan string, numRequests)

	// Создание WaitGroup
	var wg sync.WaitGroup

	// Добавление рабочих горутин в WaitGroup
	wg.Add(numRequests)

	// Запуск рабочих горутин
	for i := 0; i < numRequests; i++ {
		go worker(i+1, jobs, results, &wg)
	}

	// Запуск HTTP-сервера в тестовом режиме
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/submit" {
			t.Errorf("Expected /submit, got %s", r.URL.Path)
			return
		}
		if r.Method != "POST" {
			t.Errorf("Expected POST method, got %s", r.Method)
			return
		}
		w.WriteHeader(http.StatusOK)

		// Чтение тела запроса
		body := make([]byte, r.ContentLength)
		_, err := r.Body.Read(body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
			return
		}

		// Помещение запроса в канал задач
		jobs <- string(body)
	}))

	defer ts.Close()

	// Отправка 30 запросов
	for i := 0; i < numRequests; i++ {
		requestBody := strings.NewReader("request" + strconv.Itoa(i))
		resp, err := http.Post(ts.URL+"/submit", "text/plain", requestBody)
		if err != nil {
			t.Fatalf("Error making request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code 200, got %d", resp.StatusCode)
		}
	}

	// Закрытие канала задач и ожидание завершения всех рабочих горутин
	go func() {
		wg.Wait()
		close(results)
	}()

	// Проверка результатов
	for i := 0; i < numRequests; i++ {
		result := <-results
		t.Logf("Result: %s", result)
	}
}
