// Author: xufei
// Date: 2018/11/24

package engine

import (
	"learngo/crawler/model"
)

type Parser interface {
	Parse(contents []byte, url string) ParseResult
	SerializeParse() (name string, args ParseArgs)
}

type ParseArgs struct {
	Name   string
	Gender string
}

type ParserFunc func(content []byte, url string) ParseResult

type Request struct {
	URL    string
	Parser Parser
}

type Item struct {
	ID      string
	Payload model.Profile
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type NilParser struct{}

func (NilParser) Parse(_ []byte, _ string) ParseResult {
	return ParseResult{}
}

func (NilParser) SerializeParse() (name string, args ParseArgs) {
	return "NilParser", ParseArgs{}
}

type FuncParser struct {
	parser ParserFunc
	name   string
}

func (s *FuncParser) Parse(contents []byte, url string) ParseResult {
	return s.parser(contents, url)
}

func (s *FuncParser) SerializeParse() (name string, args ParseArgs) {
	return s.name, ParseArgs{}
}

func NewFuncParser(parser ParserFunc, name string) *FuncParser {
	return &FuncParser{
		parser: parser,
		name:   name,
	}
}
