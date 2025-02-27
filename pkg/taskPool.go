package pkg

type task struct {
	callBack func(req interface{}) interface{}
}

type taskPool struct {
	taskPool []*task
}
