package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 创建worker，等待任务被放入
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 将任务发送给调度器
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	// 从out中拿输出，存储后将后续任务发送给调度器
	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			itemCount++
			log.Printf("Got profile #%d: %v\n", itemCount, item)
			save(item)
		}

		for _, r := range result.Requests {
			if isDuplicate(r.Url) {
				continue
			}
			e.Scheduler.Submit(r)
		}
	}
}

var urls = make(map[string]bool)

func isDuplicate(url string) bool {
	if urls[url] {
		return true
	}
	urls[url] = true
	return false
}

func createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// 将该worker的channel发送给调度器
			s.WorkerReady(in)
			r := <-in
			result, e := worker(r)
			if e != nil {
				continue
			}
			out <- result
		}
	}()
}
