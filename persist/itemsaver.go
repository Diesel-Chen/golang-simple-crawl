package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() chan interface{} {
	itemChan := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-itemChan
			itemCount++
			log.Printf("itemSaver got item: #%d ,%v", itemCount, item)
			_, err := save(item)
			if err != nil {
				panic(err)
			}
		}
	}()
	return itemChan
}
func save(item interface{}) (string, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().Index("crawl").Type("zhenai").BodyJson(item).Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, nil
}
