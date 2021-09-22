package supportlib

import "sync"

type WorkerPool struct {
	wg      sync.WaitGroup
	workers int
	tasks   chan func()
}

func worker(wp *WorkerPool) {
	for task := range wp.tasks {
		task()
		wp.wg.Done()
	}
}

func CreateWorkerPool(workers int) *WorkerPool {
	wp := &WorkerPool{
		workers: workers,
		tasks:   make(chan func()),
	}
	for i := 0; i < workers; i++ {
		go worker(wp)
	}
	return wp
}

func (wp *WorkerPool) SubmitTask(task func()) {
	wp.wg.Add(1)
	wp.tasks <- task
}

func (wp *WorkerPool) WaitUntilDone() {
	wp.wg.Wait()
}
