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

func DeactivateUsersSem(usrs []users.User, wgCount int) ([]users.User, error) {
	sem := Semaphore{
		C: make(chan struct{}, wgCount),
	}

	outputCh := make(chan semResultWithError, len(usrs))
	signalCh := make(chan struct{})

	output := make([]users.User, 0, len(usrs))

	for _, v := range usrs {
		go func(usr users.User) {
			sem.Acquire()
			defer sem.Release()

			err := usr.Deactivate()

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
