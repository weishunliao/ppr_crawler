package engine

import (
	"fmt"
	"github.com/weishunliao/crawler/fetcher"
)

func Run(seeds ...Request) {
	var queue []Request
	for _, r := range seeds {
		queue = append(queue, r)
	}

	for len(queue) > 0 {
		req := queue[0]
		queue = queue[1:]

		resp, err := fetcher.Fetch(req.Url)
		if err != nil {
			fmt.Printf("Fetch Fail, err: %v\n", err)
			queue = append(queue, req)
			continue
		}
		if req.ParseFunc != nil {
			parseResult := req.ParseFunc(resp)
			queue = append(queue, parseResult.Requests...)
		}
	}
}
