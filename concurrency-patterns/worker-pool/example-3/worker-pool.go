package example_3

import (
	"fmt"
	"practice_go/concurrency-patterns/worker-pool/example-3/users"
	"sync"
)

type resultWithError struct {
	User users.User
	Err  error
}

func deactivateUser(wg *sync.WaitGroup, inCh <-chan users.User, outCh chan<- resultWithError) {
	defer wg.Done()

	for usr := range inCh {
		err := usr.Deactivate()
		outCh <- resultWithError{
			User: usr,
			Err:  err,
		}
	}
}

// DeactivateUsers Здесь же мы создаём настраиваемое заранее фиксированное количество
// горутин, которое не зависит от того, сколько задач на входе. Таким
// образом, не нужно создавать сотни тысяч горутин при соответствующем
// наборе пользователей.
// Каждая из горутин читает общий канал с входными данными, куда
// отправляются задачи, пока он не закроется.
// Аналогично предыдущему способу в основной горутине (здесь, конечно,
// речь идёт о той, в которой выполняется функция) собираются результаты.
func DeactivateUsers(usrs []users.User, wgCount int) ([]users.User, error) {
	inputCh := make(chan users.User)
	outputCh := make(chan resultWithError)
	wg := &sync.WaitGroup{}

	output := make([]users.User, 0, len(usrs))

	go func() {
		defer close(inputCh)

		for i := range usrs {
			inputCh <- usrs[i]
		}
	}()

	go func() {
		for i := 0; i < wgCount; i++ {
			wg.Add(1)

			go deactivateUser(wg, inputCh, outputCh)
		}
		wg.Wait()
		close(outputCh)
	}()

	for res := range outputCh {
		if res.Err != nil {
			return nil, fmt.Errorf("an error occurred: %w", res.Err)
		}

		output = append(output, res.User)
	}

	return output, nil
}
