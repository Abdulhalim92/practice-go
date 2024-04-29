package implementation

import "time"

type myContext struct {
}

func (c *myContext) Deadline() (deadline time.Time, ok bool) {
	return deadline, ok
}

func (c *myContext) Done() <-chan struct{} {
	myChannel := make(chan struct{})

	return myChannel
}
func (c *myContext) Err() error {
	return nil
}

func (c *myContext) Value(key interface{}) interface{} {
	return nil
}
