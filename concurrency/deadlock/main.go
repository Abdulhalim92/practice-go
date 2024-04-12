package main

import (
	"log"
)

func main() {
	//select {} // fatal error: all goroutines are asleep - deadlock!

	ch := make(chan int, 1)
	go dummy(ch)
	log.Println("waiting for reception...")
	ch <- 45
	ch <- 58
	ch <- 100
}

func dummy(c chan int) {
	smth := <-c
	log.Println("has received something", smth)
}

// 2021/02/16 11:19:57 waiting for reception...
// 2021/02/16 11:19:57 has received something 45
// fatal error: all goroutines are asleep - deadlock!

/*
	Мы создали буферизованный канал емкостью, равной 1
	Этот канал передается новой горутине, которая будет получать данные
	по каналу ( dummy функции).
	Мы отправляем по каналу три значения: 45, 58 и 100.
    Первое значение получает фиктивная горутина.
	В буфере есть место для хранения второго значения.
	Когда мы отправим третье значение на канал, основная горутина будет
	заблокирована.
	Он блокируется до тех пор, пока третье значение не будет скопировано
	в буфер канала.
	Программа будет ждать бесконечно.
*/
