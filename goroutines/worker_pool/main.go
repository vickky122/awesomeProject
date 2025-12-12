package main

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

type WorkerPool struct {
	tasks chan Task
	wg    sync.WaitGroup
}

func NewWorkerPool(workerCount int) *WorkerPool {
	p := &WorkerPool{
		tasks: make(chan Task),
	}

	for i := 0; i < workerCount; i++ {
		go p.worker(i)
	}

	return p
}

func (p *WorkerPool) worker(id int) {
	for task := range p.tasks {
		fmt.Printf("Worker %d processing task\n", id)
		task()
		p.wg.Done()
	}
}

func (p *WorkerPool) Submit(t Task) {
	p.wg.Add(1)
	p.tasks <- t
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
}

func (p *WorkerPool) Close() {
	close(p.tasks)
}

func main() {
	workerCount := 3
	totalTasks := 10

	pool := NewWorkerPool(workerCount)

	for i := 1; i <= totalTasks; i++ {
		taskNum := i
		pool.Submit(func() {
			fmt.Println("Executing task", taskNum)
			time.Sleep(500 * time.Millisecond)
		})
	}

	pool.Close()
	pool.Wait()

	fmt.Println("All tasks completed.")
}
