package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	// определение корневого контекста
	ctx1 := context.Background()
	// получение контекста с отменой
	ctx2, cancel := context.WithCancel(ctx1)
	defer cancel()
	// определение дочернего контекста с отменой
	ctx3, cancel2 := context.WithCancel(ctx2)
	defer cancel2()

	fmt.Println(reflect.TypeOf(ctx3))
}

/*
type cancelCtx struct {
	Context

	mu       sync.Mutex            // protects following fields
	done     atomic.Value          // of chan struct{}, created lazily, closed by first cancel call
	children map[canceler]struct{} // set to nil by the first cancel call
	err      error                 // set to non-nil by the first cancel call
	cause    error                 // set to non-nil by the first cancel call
}
*/

// Когда выполняем функцию отмены:
// Мьютекс будет запблкирован, никакая горуина не сможет изменить контекст.
// Канал done будет закрыт
// Все дочерные элементы будут отменены
// Мьютекс будет разблокирован
