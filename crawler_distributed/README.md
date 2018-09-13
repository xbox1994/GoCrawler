## 分布式

与并发的差别是：增加分布式去重功能、将worker分布在不同的网络节点上工作、存储服务单独开启。关键在于使用服务间调用方式将channel转换到分布式

### 运行
`go run crawler_distributed/worker/server/worker.go --port 9000`  
`go run crawler_distributed/worker/server/worker.go --port 9001`  
`go run crawler_distributed/persist/server/itemsaver.go --port 1234`  
`go run crawler_distributed/main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"`  