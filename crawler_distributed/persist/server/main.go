package main

import (
	"GoCrawler/crawler_distributed/persist"
	"GoCrawler/crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(ServeRpc(":1234", "dating_profile"))
}

func ServeRpc(host, index string) error {
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		panic(e)
	}
	rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
	return nil
}
