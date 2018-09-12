package main

import (
	"GoTest/crawler_distributed/engine"
	"GoTest/crawler_distributed/model"
	"GoTest/crawler_distributed/persist/server"
	"GoTest/crawler_distributed/rpcsupport"
	"log"
	"time"
)

func main() {
	host := ":1234"
	go server.ServeRpc(host, "test1")
	time.Sleep(time.Second)

	client, e := rpcsupport.NewClient(host)
	if e != nil {
		panic(e)
	}

	// Call save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}

	result := ""
	err := client.Call("ItemSaverService.Save", item, &result)

	if err != nil || result != "ok" {
		log.Printf("result: %s; err: %s", result, err)
	}

}
