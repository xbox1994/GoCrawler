package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
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
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		// 从out中拿输出，交给saver存储
		for _, item := range result.Items {
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		// 存储后将后续任务发送给调度器
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
