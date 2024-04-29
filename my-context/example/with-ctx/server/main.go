package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Headers: ")
		for name, headers := range r.Header {
			for _, h := range headers {
				fmt.Printf("%s: %s\n", name, h)
			}
		}

		// получение контекста запроса
		ctx := r.Context()
		// создание канала для ответа
		resChan := make(chan int)
		// запус фукции doWork в отдельном потоке
		go doWork(ctx, resChan)
		// ожидаем пока:
		// клиент не закроет соединение
		// функция doWork не завершится
		select {
		case <-ctx.Done():
			log.Println("[Handler] context cancelled in main handler, client has disconnected")
			return
		case result := <-resChan:
			log.Println("[Handler] Received 1000")
			log.Println("[Handler] Send response")
			_, err := fmt.Fprintf(w, "Response %d", result) // отправка данных клиенту
			if err != nil {
				fmt.Printf("failed to write response: %v", err)
				return
			}
			return
		}

	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicf("failed to listen and serve: %v", err)
	}
}

func doWork(ctx context.Context, resChan chan int) {
	log.Println("[doWork] launch the doWork")
	sum := 0
	for {
		log.Println("[doWork] one iteration")
		time.Sleep(time.Millisecond)
		select {
		case <-ctx.Done():
			log.Println("[doWork] ctx Done is received inside doWork")
			return
		default:
			sum++
			if sum > 1000 {
				log.Println("[doWork] sum has reached 1000")
				resChan <- sum
				return
			}
		}
	}
}
