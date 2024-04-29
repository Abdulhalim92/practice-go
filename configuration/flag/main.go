package main

import (
	"flag"
	"fmt"
)

func main() {
	// version 1 : Int
	port := flag.Int("port", 4242, "the port on which the server will listen")
	flag.Parse()
	fmt.Printf("Starting server on port %d\n", *port)

	// version 2 : IntVar
	var port2 int
	flag.IntVar(&port2, "port-2", 4242, "the port on which the server will listen")
	flag.Parse()
	fmt.Printf("Starting server-2 on port %d\n", port2)

}
