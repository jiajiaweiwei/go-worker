package main

import (
	"sync"
)

func main() {
	// 任务列表
	taskList := make([]task, 10)
	for i := 0; i < 10; i++ {
		taskList[i].callback = func() {
			println(i)
		}
	}

	// 创建协程池和任务队列
	s := server{taskList: taskList}
	s.Start(workPool{num: 3})
}

type task struct {
	callback func()
}

type workPool struct {
	num int // 工作池中协程的数量
}

type server struct {
	taskList []task
}

func (s *server) Start(workPool workPool) {
	// 任务队列
	taskQueue := make(chan task, len(s.taskList))

	// 把任务放到队列里
	for _, t := range s.taskList {
		taskQueue <- t
	}
	close(taskQueue) // 关闭任务队列，表示没有新的任务

	var wg sync.WaitGroup

	// 启动指定数量的协程处理任务
	for i := 0; i < workPool.num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range taskQueue {
				// 从队列中获取任务并执行
				task.callback()
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()
}
