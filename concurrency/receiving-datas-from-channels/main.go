package main

import "fmt"

func main() {
	var received int
	ch := make(chan int, 2)
	ch <- 42
	ch <- 41
	//ch <- 3 // произошла ошибка --- deadlock
	received = <-ch
	fmt.Println(received)

	x, ok := <-ch
	if !ok {
		fmt.Println("channel is empty or closed")
	}
	fmt.Println(x)
}
