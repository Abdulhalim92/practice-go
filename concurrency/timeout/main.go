package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan int, 1)
	select {
	case rec, ok := <-ch:
		if ok {
			log.Printf("received %d", rec)
		}
	case rec, ok := <-time.After(time.Second * 3):
		if ok {
			log.Printf("operation timed out at %s", rec)
		}
	}
}

// Output: 2024/04/09 16:40:02 operation timed out at 2024-04-09 16:40:02.4976301 +0500 +05 m=+3.002478407

// time.After возвращает канал только для приема time.Time элементов
// Через 3 секунды происходит таймаут
