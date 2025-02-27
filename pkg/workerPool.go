package pkg

import "sync"

type worker struct {
	task taskPool
	stop stopSingle
}

type workerPool struct {
	cond  sync.Cond
	works []*worker
	num   int
}

// begin working
func (w *workerPool) working() error {
	for i := 0; i < w.num; i++ {
		go func(w *workerPool) {
			select {
			case <-w.works[0].stop:

			}
		}(w)
	}
	return nil
}
