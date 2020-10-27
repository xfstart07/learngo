// Author: xufei
// Date: 2018/12/19

package service

import (
	"context"
	"learngo/crawler_distributed/config"

	"gopkg.in/olivere/elastic.v6"
)

var SearchService = &searchSrv{}

type searchSrv struct{}

var ElasticClient *elastic.Client

func NewElasticClient() error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	ElasticClient = client

	return nil
}

func (s *searchSrv) Profile(query string, page int) (*elastic.SearchHits, error) {
	q := elastic.NewQueryStringQuery(query)
	result, err := ElasticClient.
		Search(config.ElasticSearchDB).
		Type(config.ElasticSearchZhenAiTable).From(page * 10).
		Query(q).
		Do(context.Background())
	if err != nil {
		return nil, err
	}

	return result.Hits, nil
}
