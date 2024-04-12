package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Printf("Program start \n")
	// initialize the wait group
	var waitGroup sync.WaitGroup
	waitGroup.Add(10) // увеличение счетчика группы ожидания
	for i := 0; i < 10; i++ {
		go concurrentTasks(i, &waitGroup)
	}
	waitGroup.Wait() // заблокировать текущую горутину до тех пор,
	// пока все горутины не завершатся
	finishTask()
	fmt.Printf("Program end \n")
}

func finishTask() {
	fmt.Println("Executing finish task")
}

func concurrentTasks(taskNumber int, waitGroup *sync.WaitGroup) {
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)
	waitGroup.Done()
}

// Output:
// Program start
// BEGIN Execute task number 3
// BEGIN Execute task number 9
// BEGIN Execute task number 7
// ...
// END Execute task number 9
// END Execute task number 7
// END Execute task number 3
// Executing finish task
// Program end

/*
	В начале программы мы создали переменную типа sync.WaitGroup
	Мы указываем группе ожидания, что ему нужно дождаться выполнения 10
	единиц работы.
	Внутренний счетчик группы ожидания будет увеличен.
	Затем мы передаем указатель на группу ожидания нашим горутинам.
	go concurrentTasks(i, &waitGroup)
	Мы вызываем метод waitGroup.Done()в конце функции.
	Метод Done уменьшит внутренний счетчик.
	Когда внутренний счетчик достигает 0 (это означает, что все наши
	горутины завершились), основная горутина освобождается
*/
