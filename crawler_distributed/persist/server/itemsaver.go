package main

import (
	"GoCrawler/crawler_distributed/config"
	"GoCrawler/crawler_distributed/persist"
	"GoCrawler/crawler_distributed/rpcsupport"
	"flag"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(ServeRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
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
