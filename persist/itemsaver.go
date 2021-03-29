package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"golang-simple-crawl/model"
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
			err := save(item)
			if err != nil {
				panic(err)
			}
		}
	}()
	return itemChan
}
func save(item interface{}) error {
	p := item.(model.Person)
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	_, err = client.Index().Index("crawl").Type("zhenai").Id(p.Id).BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	return nil
}
