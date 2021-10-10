package engine

type SingleEngine struct {}

func (engine SingleEngine) Run(seeds ...Request) {
	var queue []Request
	for _, r := range seeds {
		queue = append(queue, r)
	}

	for len(queue) > 0 {
		req := queue[0]
		queue = queue[1:]

		parseResult, err := Execute(req)
		if err != nil {
			queue = append(queue, req)
		}
		if parseResult.Requests != nil {
			queue = append(queue, parseResult.Requests...)
		}
	}
}

