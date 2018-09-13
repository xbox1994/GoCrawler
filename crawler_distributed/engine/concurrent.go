package engine

type ConcurrentEngine struct {
	Scheduler        Scheduler
	WorkerCount      int
	ItemChan         chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParseResult, error)

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

func (e ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	// 1. 创建worker，等待任务被放入
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	// 4. 将任务发送给调度器，在调度器中将任务分配给worker
	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		// 7. 主协程一开始被阻塞，直到有任务结果被传过来
		result := <-out
		// 8. 从out中拿输出，交给saver存储
		for _, item := range result.Items {
			go func(item Item) {
				e.ItemChan <- item
			}(item)
		}

		// 9. 存储后将后续任务发送给调度器
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

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, s Scheduler) {
	go func() {
		for {
			// 2. 告诉调度器，已经有一个worker准备好了
			s.WorkerReady(in)
			// 3. 但此时没有任务塞进来，所以一开始所有worker都暂时被阻塞
			r := <-in
			// 5. 调度器将某个任务传给某个worker，执行真正的任务
			result, err := e.RequestProcessor(r)
			if err != nil {
				continue
			}
			// 6. 将任务执行结果传回engine
			out <- result
		}
	}()
}
