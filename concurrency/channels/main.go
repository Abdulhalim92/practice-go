package main

import "fmt"

// send --- chan <-int
// receive --- <-chan int
// send and receive --- chan int
func main() {
	//unbuffered()
	//buffered()
	//withGoroutine()
	withSelect()
}

func unbuffered() {
	// инициализация небуферизированных каналов
	ch1 := make(chan int)
	ch1 <- 1   // отправка данных в канал
	close(ch1) // закрытие канала
}

func buffered() {
	// инициализация буферизированных каналов
	ch2 := make(chan int, 1) // 2 - размер буфера (пропускная способность)
	ch2 <- 2                 // отправка данных в канал
	ch2 <- 3
}

func withGoroutine() {
	myChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	msg := <-myChannel

	fmt.Println(msg)
}

func withSelect() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "cow"
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
}
