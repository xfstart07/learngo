package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addPng := MakeAddSuffix(".png")

	fmt.Println(addBmp("file"))
	fmt.Println(addPng("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
