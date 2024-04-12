package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", trackVisits(homepageHandler))
	err := http.ListenAndServe(":3030", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func trackVisits(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// track the visit
		fmt.Println("one visit!")
		// call the original handler
		handler(w, r)
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my homepage")
	fmt.Fprintln(w, "I am max")
}
