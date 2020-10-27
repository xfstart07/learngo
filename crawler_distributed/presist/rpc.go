// Author: xufei
// Date: 2018/12/18

package presist

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"log"

	"gopkg.in/olivere/elastic.v6"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (s *ItemSaverService) Save(args engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, args)
	if err != nil {
		log.Printf("item saving error: %v, %v", args, err)
		return err
	}
	*result = "ok"
	return nil
}
