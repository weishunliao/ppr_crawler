package engine

import (
	"fmt"
	"github.com/weishunliao/crawler/fetcher"
)

type SingleEngine struct {}

func (engine SingleEngine) Run(seeds ...Request) {
	var queue []Request
	for _, r := range seeds {
		queue = append(queue, r)
	}

	for len(queue) > 0 {
		req := queue[0]
		queue = queue[1:]

		parseResult, err := Worker(req)
		if err != nil {
			queue = append(queue, req)
		}
		if parseResult.Requests != nil {
			queue = append(queue, parseResult.Requests...)
		}
	}
}

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
