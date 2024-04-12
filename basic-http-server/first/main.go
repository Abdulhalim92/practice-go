package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	// create a server
	myServer := &http.Server{
		// set the address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	// launch the server
	log.Fatal(myServer.ListenAndServe())
}

// sudo lsof -i 8084 -- проверка существования порта
// kill 10296 -- завершение процесса с PID 10296
// netstat -ano | find "8084" -- предоставит идентификатор процесса программы, прослушивающей порт 8084
// lsof | wc -l -- обращаемся к lsof, чтобы получить список открытых файлов, и с помощью wc -l подсчитаем количество открытых файлов
