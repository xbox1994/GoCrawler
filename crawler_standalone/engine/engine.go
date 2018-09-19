package engine

import (
	"log"
)

type SimpleEngine struct { // 创建engine结构体用以接受初始化时的参数
	WorkerCount int
}

func (e SimpleEngine) Run(seeds ...Request) { // 参数是可变长的类型，用来传入多个种子
	in := make(chan Request)      // 创建种子输入管道，其中可以传递的数据为Request，worker将从这里取，engine将种子往这里放
	out := make(chan ParseResult) // 种子输出管道，其中可以传递的数据为ParseResult，worker将解析结果传到这里交给engine

	// 开启若干个worker协程，等待种子被放入
	for i := 0; i < e.WorkerCount; i++ {
		go func() {
			for {
				// 开始时所有worker都被阻塞，等待种子被传入
				r := <-in // 创建r变量，初始值从输入管道中获取
				// 收到某个种子，开始解析
				result, e := worker(r)
				if e != nil {
					continue
				}
				// 将任务执行结果传回engine，但是这里如果直接写out <- result会无法执行，因为有循环等待产生，原因：
				// 想要执行这行代码必须需要有engine在执行39行开始等待才行，但是engine执行这行代码的前提完成48行的输入，也就是需要有空闲的worker在等待接受in，然而此时worker正在执行本行命令
				go func() { out <- result }()
			}
		}()
	}

	// 传入种子到输入管道
	for _, request := range seeds {
		in <- request
	}

	for { // 循环结构只有for关键字，这里是当成while(true)使用，让engine一直运行，等待worker返回的数据
		parseResult := <-out

		// 现在是打印得到的数据，以后将改为保存到数据库
		for _, item := range parseResult.Items { // 相当于foreach，第一个返回值是index，第二个是object，_代表不使用
			log.Printf("Got item %v\n", item)
		}

		// 存储后将后续种子发送给输入管道
		for _, r := range parseResult.Requests {
			in <- r
		}
	}
}
