// Author: xufei
// Date: 2020/3/2

package main

import (
	"fmt"
	"os"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

// NOTE: 读取的是pdf 原始内容，不能直接看
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	err := outputPdfText(inputPath)
	fmt.Println(err)
}

func outputPdfText(inputPath string) error {
	f, err := os.Open(inputPath)
	if err != nil {
		return err
	}

	defer f.Close()

	outDir := "outdir"
	if err := os.Mkdir(outDir, 0777); err != nil {
		return err
	}

	if err := pdfapi.ExtractContent(f, outDir, nil, nil); err != nil {
		return err
	}

	return nil
}
