package supportlib

import (
	"sync"
	"time"
)

func Schedule(what func(), delay time.Duration, wg *sync.WaitGroup) chan bool {
	stop := make(chan bool)
	t := time.NewTicker(delay)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-t.C:
				what()
			case <-stop:
				return
			}
		}
	}()
	return stop
}
