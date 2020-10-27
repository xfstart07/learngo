// Author: xufei
// Date: 2018/11/24

package parser

import (
	"bytes"

	"github.com/PuerkitoBio/goquery"
)

func transformDocument(body []byte) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	return doc, nil
}
