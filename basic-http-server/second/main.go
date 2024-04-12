package main

import (
	"log"
	"net/http"
	"time"
)

type myHandler struct {
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	toSend := []byte("<html><head></head><body>Hello, World!</body></html>")
	_, err := w.Write(toSend)
	if err != nil {
		log.Printf("error while writing on the body %s", err)
	}
}

func main() {
	// create a server
	myServer := &http.Server{
		// set the address
		Addr: "127.0.0.1:8080",
		// define some specific configuration
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      &myHandler{},
	}
	// launch the server
	log.Fatal(myServer.ListenAndServe())
}
