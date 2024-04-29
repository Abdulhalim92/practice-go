package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func main() {
	// создание корневого контекста
	ctx := context.Background()
	fmt.Println(reflect.TypeOf(ctx)) // *context.emptyCtx
	fmt.Println(ctx)                 // context.Background

	cancelCtx := withCancel()
	fmt.Println(reflect.TypeOf(cancelCtx))
	fmt.Println(cancelCtx)

	timeoutCtx := withTimeout()
	fmt.Println(reflect.TypeOf(timeoutCtx))
	fmt.Println(timeoutCtx)

	deadlineCtx := withDeadline()
	fmt.Println(reflect.TypeOf(deadlineCtx))
	fmt.Println(deadlineCtx)
}

// функция cancel() отменяет операцию и удаляет контекст с родительского контекста
func withCancel() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	return ctx
}

// функция withTimeout() создает контекст в таймаутом
// таймаут - максимальное количество времени, отведенное процессору для нормального завершения
func withTimeout() context.Context {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return ctx
}

// функция withDeadline() создает контекст с сроком выполнения
// deadline - определенный момент времени, когда контекст должен быть завершен
func withDeadline() context.Context {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	return ctx
}
