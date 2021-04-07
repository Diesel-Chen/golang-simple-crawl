package main

import (
	"github.com/olivere/elastic/v7"
	"golang-simple-crawl/distribute-crawl/config"
	"golang-simple-crawl/distribute-crawl/persist"
	"golang-simple-crawl/distribute-crawl/rpcsupport"
	"log"
)

func main() {
	log.Fatal(ItemSaverRpc(config.ItemSaverHost, config.ElasticIndex))
}
func ItemSaverRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	err = rpcsupport.ServeRpc(host,
		&persist.ItemSaverRPC{
			Esclient: client,
			Index:    index,
		})
	if err != nil {
		return err
	}
	return nil

}
