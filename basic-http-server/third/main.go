package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := os.Open("./basic-http-server/third/tmp/test.txt")
	if err != nil {
		log.Printf("error file opening %s", err)
	} else {
		log.Println("file opened")
	}
	// всегда нужно закрыть файл, чтобы не засорять память
}

func main() {
	// create a server
	myServer := http.Server{
		// set the server address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &myHandler{},
	}

	// launch the server
	log.Fatal(myServer.ListenAndServe())
}

// lsof -i :8080 -- поиск идентификатора процесса, прослушивающий порт 8080
// lsof -p 12458 | wc -l -- подсчет поличество открытых файлов с указанным идентификатором процесса
