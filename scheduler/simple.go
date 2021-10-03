package scheduler

import "github.com/weishunliao/crawler/engine"

type SimpleScheduler struct {
	WorkerChannel chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChannel(ch chan engine.Request) {
	s.WorkerChannel = ch
}

func (s *SimpleScheduler) Submit(req engine.Request) {
	go func() {
		s.WorkerChannel <- req
	}()
}


