package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int)
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	log.Println("received")
}

func dummy(c chan int) {
	time.Sleep(3 * time.Second)
	<-c
}

// Когда вы отправляете данные в небуферизованный канал, ваша текущая
// горутина будет заблокирована до тех пор, пока данные не будут
// получены другой горутиной.

// 2024/04/09 15:05:42 waiting for reception...
// 2024/04/09 15:05:45 received
