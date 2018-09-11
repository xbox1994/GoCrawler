package scheduler

import "GoTest/crawler_distributed/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var rQ []engine.Request
		var wQ []chan engine.Request
		for {
			var activeR engine.Request
			var activeW chan engine.Request
			if len(rQ) > 0 && len(wQ) > 0 {
				activeW = wQ[0]
				activeR = rQ[0]
			}
			select {
			case r := <-s.requestChan:
				rQ = append(rQ, r)
			case w := <-s.workerChan:
				wQ = append(wQ, w)
			case activeW <- activeR:
				wQ = wQ[1:]
				rQ = rQ[1:]
			}
		}
	}()
}
