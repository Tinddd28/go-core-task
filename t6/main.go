package t7

import (
	"math/rand"
	"sync"
	"time"
)

func RandGenerator() (chan int, func()) {
	ch := make(chan int)
	stop := make(chan struct{})
	var once sync.Once

	go func() {
		rand.Seed(time.Now().UnixNano())
		for {
			select {
			case <-stop:
				close(ch)
				return
			case ch <- rand.Int():
			}
		}
	}()

	stopFunc := func() {
		once.Do(func() {
			close(stop)
		})
	}
	return ch, stopFunc
}
