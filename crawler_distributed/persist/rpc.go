package persist

import (
	"GoCrawler/crawler_distributed/config"
	"GoCrawler/crawler_distributed/engine"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	e := Save(s.Client, item, s.Index)
	if e == nil {
		*result = "ok"
		log.Printf("Success saving %v", item)
	} else {
		log.Printf("Error saving %v, %v", item, e)
	}
	return e
}

func Save(client *elastic.Client, item engine.Item, index string) error {
	_, e := client.Index().Index(index).Type(config.ElasticType).Id(item.Id).BodyJson(item).Do(context.Background())
	if e != nil {
		log.Printf("es create index fail %v", e)
		return e
	}
	return nil
}
