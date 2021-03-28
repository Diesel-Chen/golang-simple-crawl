package persist

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"golang-simple-crawl/model"
	"testing"
)

func TestSave(t *testing.T) {
	excepted := model.Person{
		Name:           "test",
		Age:            18,
		Height:         "160cm",
		Weight:         "60kg",
		Constellation:  "处女座",
		Salary:         "10000-20000",
		Job:            "老师",
		Degree:         "本科",
		WorkStation:    "上海",
		MarriageStatus: "未婚",
	}
	id, err := save(excepted)
	if err != nil {
		panic(err)
	}
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	resp, err := client.Get().Index("crawl").Type("zhenai").Id(id).Do(context.Background())
	var actual model.Person
	json.Unmarshal(resp.Source, &actual)
	if excepted != excepted {
		t.Errorf("expected %v got:%v", excepted, actual)
	}
}
