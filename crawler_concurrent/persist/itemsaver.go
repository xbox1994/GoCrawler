package persist

import (
	"GoCrawler/crawler_concurrent/engine"
	"context"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, e := elastic.NewClient(elastic.SetSniff(false))
	if e != nil {
		log.Printf("es client create fail %v", e)
		return nil, e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item saver got profile #%d: %v\n", itemCount, item)

			e := save(client, item, index)
			if e != nil {
				log.Printf("Item saver save item error: %v %v\n", item, e)
				continue
			}

		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {
	_, e := client.Index().Index(index).Type("zhenai").Id(item.Id).BodyJson(item).Do(context.Background())
	if e != nil {
		log.Printf("es create index fail %v", e)
		return e
	}
	return nil

}
