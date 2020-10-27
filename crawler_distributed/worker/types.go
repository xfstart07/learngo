// Author: xufei
// Date: 2018/12/18

package worker

import (
	"fmt"
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler_distributed/config"
	"log"
)

type SerializedParser struct {
	Name string
	Args engine.ParseArgs
}

type Request struct {
	Url    string
	Parser SerializedParser
}

func SerializeRequest(r engine.Request) Request {
	name, args := r.Parser.SerializeParse()

	return Request{
		Url: r.URL,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func DeSerializeRequest(r Request) (engine.Request, error) {
	switch r.Parser.Name {
	case config.ProfileParser:
		return engine.Request{
			URL:    r.Url,
			Parser: parser.NewProfileParser(r.Parser.Args),
		}, nil
	case config.ParseCity:
		return engine.Request{
			URL:    r.Url,
			Parser: engine.NewFuncParser(parser.ParseCity, config.ParseCity),
		}, nil
	case config.ParseCityList:
		return engine.Request{
			URL:    r.Url,
			Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
		}, nil
	default:
		return engine.Request{}, fmt.Errorf("error parser name")
	}
}

type ParseResult struct {
	Items    []engine.Item
	Requests []Request
}

func SerializeResult(result engine.ParseResult) ParseResult {
	parseResult := ParseResult{
		Items: result.Items,
	}

	for _, req := range result.Requests {
		parseResult.Requests = append(parseResult.Requests, SerializeRequest(req))
	}

	return parseResult
}

func DeserializeResult(r ParseResult) engine.ParseResult {
	result := engine.ParseResult{
		Items: r.Items,
	}

	for _, req := range r.Requests {
		request, err := DeSerializeRequest(req)
		if err != nil {
			log.Printf("request err %v, %v", req, err)
		}
		result.Requests = append(result.Requests, request)
	}

	return result
}
