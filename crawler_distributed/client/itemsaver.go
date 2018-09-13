package client

import (
	"GoCrawler/crawler_distributed/engine"
	"GoCrawler/crawler_distributed/rpcsupport"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, e := rpcsupport.NewClient(host)
	if e != nil {
		return nil, e
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			log.Printf("Item saver got profile #%d: %v\n", itemCount, item)

			result := ""
			client.Call("ItemSaverService.Save", item, &result)
			if e != nil {
				log.Printf("Item saver save item error: %v %v\n", item, e)
				continue
			}

		}
	}()
	return out, nil
}
