package engine

import "time"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(Request)
	ConfigMasterWorkerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

var t = time.Tick(10 * time.Millisecond)

func (e *ConcurrentEngine) Run(seed ...Request) {
	<-t
	out := make(chan RequestResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(out, e.Scheduler)
	}

	for _, v := range seed {
		if isDuplicate(v.Url) {
			continue
		}
		e.Scheduler.Submit(v)
	}

	for {
		result := <-out
		for _, v := range result.Items {
			e.ItemChan <- v
		}
		for _, v := range result.Requests {
			if isDuplicate(v.Url) {
				continue
			}
			e.Scheduler.Submit(v)
		}
	}
}

func (e *ConcurrentEngine) createWorker(out chan RequestResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			result := Work(request)
			out <- result
		}
	}()
}

var dupMap = make(map[string]bool)

func isDuplicate(url string) bool {
	if dupMap[url] {
		return true
	}
	dupMap[url] = true
	return false

}
