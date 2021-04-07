package client

import (
	"golang-simple-crawl/crawl/model"
	"golang-simple-crawl/distribute-crawl/config"
	"golang-simple-crawl/distribute-crawl/rpcsupport"
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

			//err := Save(item)
			//CallRpc
			client, err := rpcsupport.NewClient(config.ItemSaverHost)
			if err != nil {
				panic(err)
			}
			p := item.(model.Person)
			result := ""
			err = client.Call(config.ItemSaverRPC, p, &result)
			if err != nil {
				log.Printf("Item saver error . got item %v err %s", p, err)
			}
		}
	}()
	return itemChan
}
