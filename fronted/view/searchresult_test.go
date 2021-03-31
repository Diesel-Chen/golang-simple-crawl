package view

import (
	"golang-simple-crawl/fronted/model"
	model2 "golang-simple-crawl/model"
	"log"
	"os"
	"testing"
)

func TestSearchResultView_Render(t *testing.T) {
	SearchResultView := CreateSearchResultView("template.html")
	file, err := os.Create("template_test.html")
	if err != nil {
		panic(err)
	}
	item := make([]interface{}, 0)
	m := model2.Person{
		Id:             "123123",
		Url:            "http://test/123123",
		Name:           "test",
		Age:            18,
		Height:         160,
		Weight:         60,
		Constellation:  "处女座",
		Salary:         "10000-20000",
		Job:            "老师",
		Degree:         "本科",
		WorkStation:    "上海",
		MarriageStatus: "未婚",
	}
	for i := 0; i < 10; i++ {
		item = append(item, m)
	}
	page := model.SearchResult{
		Hits:  100,
		Start: 0,
		Items: item,
	}
	log.Fatal(SearchResultView.Render(file, page))
}
