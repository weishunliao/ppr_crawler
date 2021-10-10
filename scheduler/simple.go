package scheduler

import "github.com/weishunliao/crawler/engine"

type SimpleScheduler struct {
	WorkerChannel chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.WorkerChannel
}

func (s *SimpleScheduler) WorkerReady(r chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.WorkerChannel = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	go func() {
		s.WorkerChannel <- req
	}()
}


