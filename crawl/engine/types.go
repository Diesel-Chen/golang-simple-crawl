package engine

type ParserFunc func([]byte) RequestResult

type Request struct {
	Url        string
	ParserFunc ParserFunc
}

type RequestResult struct {
	Requests []Request
	Items    []interface{}
}

func NilParserFunc(cnt []byte) RequestResult {
	return RequestResult{}
}
