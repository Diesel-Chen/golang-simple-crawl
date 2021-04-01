package model

type SearchResult struct {
	Hits     int
	Start    int
	Query    string
	PreFrom  int
	NextFrom int
	Items    []interface{}
}
