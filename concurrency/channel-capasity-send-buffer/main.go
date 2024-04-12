package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	log.Println("received")
}

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}

// Когда вы отправляете данные в буферизованный канал, ваша текущая
// горутина будет заблокирована до тех пор, пока данные не будут
// скопированы в буфер

// 2024/04/09 15:05:04 waiting for reception...
// 2024/04/09 15:05:04 received
