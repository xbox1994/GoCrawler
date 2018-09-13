package persist

import (
	"GoCrawler/crawler_distributed/engine"
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
