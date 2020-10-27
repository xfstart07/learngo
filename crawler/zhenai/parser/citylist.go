// Author: xufei
// Date: 2018/11/23

package parser

import (
	"fmt"
	"learngo/crawler/engine"
	"log"

	"github.com/PuerkitoBio/goquery"
)

const (
	CitySelector = "article.g-container > .city-list a"
)

type City struct {
	Name string
	URL  string
}

func ParseCityList(body []byte, _ string) engine.ParseResult {
	doc, err := transformDocument(body)
	if err != nil {
		log.Printf("transform document error: %v", err)
		return engine.ParseResult{}
	}

	var result engine.ParseResult
	selections := doc.Find(CitySelector)

	selections.Each(func(i int, sel *goquery.Selection) {
		url, ok := sel.Attr("href")
		if !ok {
			log.Printf("href attr not ok")
			return
		}
		log.Printf("index=%d, text=%s, url=%s\n", i, sel.Text(), url)

		result.Requests = append(result.Requests, engine.Request{
			URL:    url,
			Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
		})

		// FIXME: 获取20页的用户数据
		for i := 2; i < 20; i++ {
			result.Requests = append(result.Requests, engine.Request{
				URL:    fmt.Sprintf("%s/%d", url, i),
				Parser: engine.NewFuncParser(ParseCity, "ParseCity"),
			})
		}

	})

	return result
}
