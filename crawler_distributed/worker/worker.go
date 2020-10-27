// Author: xufei
// Date: 2018/12/18

package worker

import (
	"learngo/crawler/engine"
)

type CrawlService struct{}

func (s *CrawlService) Process(req *Request, result *ParseResult) error {
	request, err := DeSerializeRequest(*req)
	if err != nil {
		return err
	}

	parseResult, err := engine.Worker(request)
	if err != nil {
		return err
	}

	*result = SerializeResult(parseResult)

	return nil
}
