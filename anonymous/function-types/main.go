package main

import "log"

type Funky func(string)

func main() {
	var f Funky
	f = func(s string) {
		// my function defined
		log.Printf("Funky: %s", s)
	}

	f("Groovy")
}
