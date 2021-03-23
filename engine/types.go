package engine

type Request struct{
	Url string
	ParserFunc func([]byte)RequestResult
}

type RequestResult struct{
	Requests []Request
	Items []interface{}
}

func NilParserFunc(cnt []byte)RequestResult{
	return RequestResult{}
}
