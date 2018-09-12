package persist

import (
	"GoTest/crawler_distributed/engine"
	"gopkg.in/olivere/elastic.v5"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	e := Save(s.Client, item, s.Index)
	if e == nil {
		*result = "ok"
	}
	return e
}
