package main

import (
	"GoCrawler/crawler_distributed/client"
	"GoCrawler/crawler_distributed/engine"
	"GoCrawler/crawler_distributed/parser"
	"GoCrawler/crawler_distributed/scheduler"
)

func main() {
	itemChan, e := client.ItemSaver(":1234")
	if e != nil {
		panic(e)
	}

	engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    itemChan,
	}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//
	//host := ":1234"
	//go server.ServeRpc(host, "test1")
	//time.Sleep(time.Second)
	//
	//client, e := rpcsupport.NewClient(host)
	//if e != nil {
	//	panic(e)
	//}
	//
	//// Call save
	//item := engine.Item{
	//	Url:  "http://album.zhenai.com/u/108906739",
	//	Type: "zhenai",
	//	Id:   "108906739",
	//	Payload: model.Profile{
	//		Age:        34,
	//		Height:     162,
	//		Weight:     57,
	//		Income:     "3001-5000元",
	//		Gender:     "女",
	//		Name:       "安静的雪",
	//		Xinzuo:     "牡羊座",
	//		Occupation: "人事/行政",
	//		Marriage:   "离异",
	//		House:      "已购房",
	//		Hokou:      "山东菏泽",
	//		Education:  "大学本科",
	//		Car:        "未购车",
	//	},
	//}
	//
	//result := ""
	//err := client.Call("ItemSaverService.Save", item, &result)
	//
	//if err != nil || result != "ok" {
	//	log.Printf("result: %s; err: %s", result, err)
	//}

}
