package persist

import "log"

func ItemSaver() chan interface{} {
	itemChan := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-itemChan
			itemCount++
			log.Printf("itemSaver got item: #%d ,%s", itemCount, item)
		}
	}()
	return itemChan
}
