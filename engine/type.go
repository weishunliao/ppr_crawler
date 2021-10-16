package engine

type ParseResult struct {
	Requests   []Request
	Properties []Property
	Store      bool
}

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type Property interface {}
