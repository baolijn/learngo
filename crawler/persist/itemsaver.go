package persist

import (
	"log"
	"gopkg.in/olivere/elastic.v5"
	"golang.org/x/net/context"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out
}

func save(item interface{})(id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://139.219.99.228:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		return "", nil
	}

	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).Do(context.Background())

	if err != nil {
		return "", nil
	}

	return resp.Id, nil
}
