package main

import (
	"GoCrawler/crawler_distributed/config"
	"GoCrawler/crawler_distributed/persist"
	"GoCrawler/crawler_distributed/rpcsupport"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func main() {
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", config.ItemSaverPort), config.ElasticIndex))
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
