package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Program start \n")
	for i := 0; i < 10; i++ {
		go concurrentTasks(i)
	}
	finishTask()
	fmt.Printf("Program end\n")
}

func finishTask() {
	fmt.Printf("Executing finish task")
}

func concurrentTasks(taskNumber int) {
	fmt.Printf("BEGIN Execute task number %d \n", taskNumber)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("END Execute task number %d \n", taskNumber)
}

// Горутины не выполнились...Почему?
// Потому что, запуск горутины не блокирует основную горутину
// Program start
// Executing finish taskProgram end
