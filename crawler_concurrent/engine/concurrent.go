package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureWorkerChan(chan Request)
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)
	e.Scheduler.ConfigureWorkerChan(in)

	// 创建worker等待任务放入in
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
	}

	// 将任务放入in
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	// 主协程打印
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got Item: %v\n", item)
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			r := <-in
			result, e := worker(r)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
