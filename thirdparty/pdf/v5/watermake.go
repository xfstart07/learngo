// Author: xufei
// Date: 2020/3/2

package main

import (
	"fmt"
	"os"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	pdfapi "github.com/pdfcpu/pdfcpu/pkg/api"
)

// NOTE: 添加水印
func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run pdf_extract_text.go input.pdf\n")
		os.Exit(1)
	}
	inputPath := os.Args[1]

	if err := watermake(inputPath); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("水印添加成功")
	}
}

// outputPdfText prints out contents of PDF file to stdout.
func watermake(inFile string) error {
	desc := "f:Helvetica, s:0.5, rot:20"
	onTop := false
	modeParam := "LearnGo"
	selectedPages := []string{"1-"}

	wm, err := pdfcpu.ParseTextWatermarkDetails(modeParam, desc, onTop)
	if err != nil {
		return err
	}

	outFile := "output.pdf"
	if err := pdfapi.AddWatermarksFile(inFile, outFile, selectedPages, wm, nil); err != nil {
		return fmt.Errorf("%s: %w", outFile, err)
	}

	if err := pdfapi.ValidateFile(outFile, nil); err != nil {
		return fmt.Errorf("watermark: %s", err)
	}

	return nil
}
