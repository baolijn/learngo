package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
	"learngo/crawler/engine"
	"github.com/pkg/errors"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://139.219.99.228:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		return  nil
	}

	if item.Type == "" {
		return errors.New("Must supply Type")
	}

	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)

		if item.Id != "" {
			indexService.Id(item.Id)
		}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return nil
	}

	return  nil
}
