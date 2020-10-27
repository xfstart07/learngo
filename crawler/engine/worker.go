// Author: xufei
// Date: 2018/12/18

package engine

import (
	"learngo/crawler/fetcher"
	"log"
)

func Worker(request Request) (ParseResult, error) {
	body, err := fetcher.Fetch(request.URL)
	if err != nil {
		log.Printf("Fetcher: error fetching %v", err)
		return ParseResult{}, err
	}

	return request.Parser.Parse(body, request.URL), nil
}
