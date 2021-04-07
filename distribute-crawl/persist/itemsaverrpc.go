package persist

import (
	"context"
	"github.com/olivere/elastic/v7"
	"golang-simple-crawl/crawl/model"
	"log"
)

type ItemSaverRPC struct {
	Esclient *elastic.Client
	Index    string
}

func (i *ItemSaverRPC) Save(item model.Person, result *string) error {
	log.Printf("RPC item save %v", item)
	_, err := i.Esclient.Index().Index(i.Index).Type("zhenai").Id(item.Id).BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}
	*result = "ok"
	return nil
}
