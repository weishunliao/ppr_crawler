package engine

type ParseResult struct {
	Requests   []Request
	Properties []interface{}
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

func NilParser([] byte) ParseResult {
	return ParseResult{}
}
