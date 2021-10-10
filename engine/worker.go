package engine

import (
	"fmt"
	"github.com/weishunliao/crawler/fetcher"
)

func Worker(req Request) (ParseResult, error) {
	resp, err := fetcher.Fetch(req.Url)
	if err != nil {
		fmt.Printf("Fetch Fail, err: %v\n", err)
		return ParseResult{}, err
	}
	if req.ParseFunc != nil {
		return req.ParseFunc(resp), nil
	}
	return ParseResult{}, nil
}
