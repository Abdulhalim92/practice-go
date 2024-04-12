package main

// send --- chan <-int
// receive --- <-chan int
// send and receive --- chan int
func main() {
	// инициализация небуферизированных каналов
	ch1 := make(chan int)
	ch1 <- 1   // отправка данных в канал
	close(ch1) // закрытие канала

	// инициализация буферизированных каналов
	ch2 := make(chan int, 2) // 2 - размер буфера (пропускная способность)
	ch2 <- 2                 // отправка данных в канал
}
