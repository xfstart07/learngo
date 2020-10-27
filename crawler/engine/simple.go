// Author: xufei
// Date: 2018/11/24

package engine

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, seed := range seeds {
		requests = append(requests, seed)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		for _, r := range parseResult.Requests {
			requests = append(requests, r)
		}
	}

}
