package pattern

import (
	"fmt"
	"practice_go/concurrency-patterns/worker-pool/pattern/users"
)

type Semaphore struct {
	C chan struct{}
}

func (s *Semaphore) Acquire() {
	s.C <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.C
}

type semResultWithError struct {
	User users.User
	Err  error
}

// DeactivateUsersSem В рамках данного подхода создаётся, по сути, неограниченное количество
// горутин. Их выполнение блокируется с помощью объекта Semaphore, который
// ограничивает количество одновременно выполняемых горутин с помощью
// буферизированного канала. В данном подходе на каждую задачу создаётся
// горутина. Однако если буфер канала семафора переполнен, то при операции
// Acquire горутина с задачей блокируется до тех пор, пока буфер не освободится
// с помощью операции Release.
// Для сбора результатов используется канал, из которого происходит чтение
// в основной горутине (здесь, конечно, речь идёт о той, в которой выполняется
// функция, это совсем не обязательно может быть горутина, в которой выполняется
// функция main).
func DeactivateUsersSem(usrs []users.User, wgCount int) ([]users.User, error) {
	// создаем семафор и передаем ему канал с размером буфера равным
	// ограничению на количество одновременно выполняемых горутин
	sem := Semaphore{
		C: make(chan struct{}, wgCount),
	}

	// канал для сбора результатов
	outputCh := make(chan semResultWithError, len(usrs))
	// канал для оповещения горутин, что мы больше не ждем их выполнения
	signalCh := make(chan struct{})

	output := make([]users.User, 0, len(usrs))

	for _, v := range usrs {
		go func(usr users.User) {
			sem.Acquire()
			defer sem.Release()

			err := usr.Deactivate()

			// если ловим закрытие сигнального канала, то завершаем функцию
			select {
			case outputCh <- semResultWithError{
				User: usr,
				Err:  err,
			}:
			case <-signalCh:
				return
			}
		}(v)
	}

	// ждем и собираем результаты либо мы получим все, либо выйдет ошибка,
	// по которой мы перестанем читать
	for i := len(usrs); i > 0; i-- {
		res := <-outputCh
		if res.Err != nil {
			close(signalCh)
			return nil, fmt.Errorf("an error occured: %w", res.Err)
		}

		output = append(output, res.User)
	}

	return output, nil
}
