// Author: xufei
// Date: 2020/3/3

package main

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

var (
	inDir = "../../testdata"
)

func TestExtractContentCommand(t *testing.T) {
	msg := "TestExtractContentCommand"
	selectPage := []string{"1"}
	// Extract content of all pages into outDir.
	inFile := filepath.Join(inDir, "input.pdf")
	fmt.Println(inFile)
	outDir := "outdir"
	if err := api.ExtractContentFile(inFile, outDir, selectPage, pdfcpu.NewDefaultConfiguration()); err != nil {
		t.Fatalf("%s %s: %v\n", msg, inFile, err)
	}
}
