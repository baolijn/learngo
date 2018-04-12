package main

import (
	"gopkg.in/olivere/elastic.v5"
	"learngo/crawler_distributed/rpcsupport"
	"learngo/crawler_distributed/persist"
	"log"
)

func main() {
	log.Fatal(serveRpc(":1234", "dating_profile"))
}

func serveRpc(host string, Index string) error {
	client, err := elastic.NewClient(
		elastic.SetURL("http://139.219.99.228:9200/"),
		elastic.SetSniff(false))
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index: Index,
	})
}