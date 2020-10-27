// Author: xufei
// Date: 2018/11/22

package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}

	selections := doc.Find("article.g-container > .city-list a")
	fmt.Printf("len = %d\n", len(selections.Nodes))

	selections.Each(func(i int, sel *goquery.Selection) {
		url, _ := sel.Attr("href")
		fmt.Printf("index=%d, text=%s, url=%s\n", i, sel.Text(), url)
	})
}
