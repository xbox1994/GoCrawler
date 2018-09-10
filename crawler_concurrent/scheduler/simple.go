package scheduler

import "GoTest/crawler_concurrent/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) ConfigureWorkerChan(c chan engine.Request) {
	s.workerChan = c
}
