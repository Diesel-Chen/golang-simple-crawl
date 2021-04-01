package scheduler

import "golang-simple-crawl/crawl/engine"

type SimpleScheduler struct {
	WorkerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigMasterWorkerChan(requests chan engine.Request) {
	s.WorkerChan = requests
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.WorkerChan <- request
	}()

}
