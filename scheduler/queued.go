package scheduler

import "golang-simple-crawl/engine"

type QueuedScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.RequestChan <- request
}

func (s *QueuedScheduler) WorkerReady(r chan engine.Request) {
	s.WorkerChan <- r
}

func (s *QueuedScheduler) ConfigMasterWorkerChan(requests chan engine.Request) {
	panic("implement me")
}

func (s *QueuedScheduler) Run() {
	s.RequestChan = make(chan engine.Request)
	s.WorkerChan = make(chan chan engine.Request)
	go func() {
		var RequestQ []engine.Request
		var WorkerQ []chan engine.Request
		for {
			var ActiveRequest engine.Request
			var ActiveWorker chan engine.Request
			if len(RequestQ) > 0 && len(WorkerQ) > 0 {
				ActiveRequest = RequestQ[0]
				ActiveWorker = WorkerQ[0]
			}
			select {
			case r := <-s.RequestChan:
				RequestQ = append(RequestQ, r)
			case w := <-s.WorkerChan:
				WorkerQ = append(WorkerQ, w)
			case ActiveWorker <- ActiveRequest:
				RequestQ = RequestQ[1:]
				WorkerQ = WorkerQ[1:]
			}
		}
	}()
}
