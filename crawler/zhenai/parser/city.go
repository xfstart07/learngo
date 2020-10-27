// Author: xufei
// Date: 2018/11/23

package parser

import (
	"learngo/crawler/engine"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const (
	UserSelector = ".g-list .list-item"
)

var (
	profileUniqMap = make(map[string]bool)
	profileMtx     = sync.Mutex{}
)

func ParseCity(body []byte, _ string) engine.ParseResult {
	doc, err := transformDocument(body)
	if err != nil {
		log.Printf("transform document error: %v", err)
		return engine.ParseResult{}
	}

	var result engine.ParseResult
	doc.Find(UserSelector).Each(func(index int, selection *goquery.Selection) {
		url, ok := selection.Find(".photo a").Attr("href")
		if !ok {
			log.Println("无用户地址")
			return
		}

		// 去重
		profileMtx.Lock()
		defer profileMtx.Unlock()
		if profileUniqMap[url] {
			log.Printf("重复URL %s", url)
			return
		}
		profileUniqMap[url] = true

		name := selection.Find("tr:first-child").Text()
		gender := selection.Find("tr:nth-child(2) td:first-child").Text()
		gender = strings.Replace(gender, "性别：", "", -1)

		result.Requests = append(result.Requests, engine.Request{
			URL:    url,
			Parser: NewProfileParser(engine.ParseArgs{Name: name, Gender: gender}),
		})

	})

	return result
}
