package controller

import (
	"context"
	"golang-simple-crawl/crawl/fronted/model"
	"golang-simple-crawl/crawl/fronted/view"
	model2 "golang-simple-crawl/crawl/model"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
)

type SearchResult struct {
	View     view.SearchResultView
	EsClient *elastic.Client
}

func CreateSerachResult() SearchResult {
	view := view.CreateSearchResultView("fronted/view/template.html")
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResult{
		View:     view,
		EsClient: client,
	}

}

func (s SearchResult) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}
	result, err := s.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.View.Render(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
func (s SearchResult) getSearchResult(query string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	resp, err := s.EsClient.Search("crawl").Query(elastic.NewQueryStringQuery(query)).From(from).Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = int(resp.TotalHits())
	result.Start = from
	result.Query = query
	result.Items = resp.Each(reflect.TypeOf(model2.Person{}))
	result.PreFrom = from - 10
	result.NextFrom = from + 10
	return result, nil
}
