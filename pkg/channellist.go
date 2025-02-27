package pkg

type singleChan chan struct{}

type stopSingle chan struct{}

type channelPool struct {
	taskPool taskPool
	single   singleChan
}
