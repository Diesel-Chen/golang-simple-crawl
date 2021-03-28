package persist

import (
	"bytes"
	"fmt"
	"golang-simple-crawl/model"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func PictureSaver() chan interface{} {
	itemChan := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-itemChan
			itemCount++
			log.Printf("itemSaver got item: #%d ,%v", itemCount, item)
			err := SavePicture(item)
			if err != nil {
				panic(err)
			}
		}
	}()
	return itemChan
}

func SavePicture(item interface{}) error {
	u := item.(model.Picture)
	storePath := "/Users/dieselchen/Pictures/crawl/"
	resp, err := http.Get(u.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	name := fmt.Sprintf("[%s]%d%s", u.Name, time.Now().UnixNano(), path.Base(u.Url))
	file, err := os.Create(storePath + name)
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := io.Copy(file, bytes.NewBuffer(body))
	fmt.Println("n:", n)
	return nil
}
