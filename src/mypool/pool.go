package mypool

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	OutChan  chan *Task
	jobsChan chan *Task // job in work
	Cap      int        // the num of worker
}

// NewPool new a new mypool
func NewPool(cap int) *Pool {
	return &Pool{
		OutChan:  make(chan *Task),
		jobsChan: make(chan *Task),
		Cap:      cap,
	}
}

// make a goroutine
func (p *Pool) newWorker(i int) {
	for task := range p.jobsChan {
		task.execute()
		time.Sleep(1 * time.Second)
		fmt.Println("Worker:", i, "finish.")
	}
}

// Run pool
func (p *Pool) Run() {
	for i := 0; i < p.Cap; i++ {
		go p.newWorker(i)
	}
	for task := range p.OutChan {
		p.jobsChan <- task
	}
}

// 提高 waitGroup的易用性
type WarpGroup struct {
	waitGroup sync.WaitGroup
}

func (w *WarpGroup) Warp(f func()) {
	w.waitGroup.Add(1)
	go func() {
		defer w.waitGroup.Done()
		f()
	}()
}

func (w *WarpGroup) Exist() {
	w.waitGroup.Wait()
}
