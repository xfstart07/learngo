// Author: xufei
// Date: 2020/3/2

package main

import (
	"bytes"
	"fmt"

	"github.com/ledongthuc/pdf"
)

// NOTE: 读取的内容没有换行
func main() {
	//content, err := readPdf("input.pdf") // Read local pdf file
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(content)
	//
	//pageContent, err := readPdfPageContent("input.pdf")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(pageContent)

	content, err := readPdfByPage("input.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
}

// 读取信息
func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	if _, err := buf.ReadFrom(b); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// 读取其中一页的文字内容
func readPdfPageContent(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	b, err := r.Page(7).GetPlainText(nil)
	if err != nil {
		return "", err
	}

	return b, nil
}

// 一页一页读
func readPdfByPage(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	var content string
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		text, err := p.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		fmt.Println(text)

		content += text + "\n"
	}

	return content, nil
}
