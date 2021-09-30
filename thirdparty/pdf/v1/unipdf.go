// Author: xufei
// Date: 2020/3/2

package main

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	pdf "github.com/unidoc/unipdf/v3/model"
)

// 提取 PDF 文字，unipdf 库，没有证书
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	err := outputPdfText(inputPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	pdfReader, err := pdf.NewPdfReader(f)
	if err != nil {
		return err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return err
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			return err
		}

		ex, err := extractor.New(page)
		if err != nil {
			return err
		}

		text, err := ex.ExtractText()
		if err != nil {
			return err
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}

	return nil
}
