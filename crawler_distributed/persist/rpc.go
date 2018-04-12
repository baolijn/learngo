package persist

import (
	"learngo/crawler/engine"
	"learngo/crawler/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index string
}

func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)

	if err == nil {
		*result = "ok"
	}else{
		log.Printf("Error saving item %v: %v", item, err)
	}

	return err
}
