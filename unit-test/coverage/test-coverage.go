package coverage

import (
	"fmt"
	"sync"
)

func BazBaz(number int) int {
	var waitGroup sync.WaitGroup
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go concurrentTask(number, &waitGroup)
	}
	waitGroup.Wait()
	return number
}

func concurrentTask(number int, waitGroup *sync.WaitGroup) {
	useless := number + 2
	fmt.Println(useless)
	waitGroup.Done()
}
