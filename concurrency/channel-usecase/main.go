package main

import (
	"fmt"
	"time"
)

func main() {
	syncCh := make(chan bool)
	// launch a second goroutine
	go func() {
		longTask2()
		// finished
		syncCh <- true
	}()
	longTask()
	// blocks until the second goroutine finishes
	<-syncCh
}

// Небуферизованные каналы используются для синхронизации двух горутин

func longTask() {
	time.Sleep(3 * time.Second)
	fmt.Println("long task finished")
}

func longTask2() {
	time.Sleep(1 * time.Second)
	fmt.Println("long task 2 finished")
}

/*
	Небуферизованный канал используется здесь для синхронизации основной
	горутины со второй горутиной. Операция приема <-syncCh блокируется
	до тех пор, пока не завершится другая горутина. Чтобы сигнализировать
	по завершении, вторая горутина отправит в канал значение «true».
*/
