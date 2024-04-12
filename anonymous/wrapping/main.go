package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/homepage", homepageHandler)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to my homepage")
	fmt.Fprintln(w, "I am max")
}
