package engine

type ParseResult struct {
	Requests   []Request
	Properties []interface{}
	Store      bool
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}
