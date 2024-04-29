package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// В функции нужно запустить несколько горутин с вычислениями и дождаться их
// окончания.
// При этом, если функция будет работать больше указанного количества секунд,
// нужно прервать ее выполнение.

func main() {
	timeLimit := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeLimit)
	defer cancel()

	fun(ctx)
}

func isDone(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func FibonacciRecursion(ctx context.Context, n int) int {
	if n <= 1 || isDone(ctx) {
		return n
	}
	return FibonacciRecursion(ctx, n-1) + FibonacciRecursion(ctx, n-2)
}

func f(ctx context.Context, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(FibonacciRecursion(ctx, 1+i))
}

func fun(ctx context.Context) {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go f(ctx, i, &wg)
	}
	wg.Wait()
	fmt.Println("Function ended")
}
