package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler/scheduler"
	"learngo/crawler_distributed/persist/client"
	"fmt"
	"learngo/crawler_distributed/config"
)

func main() {
	itemChan, err := client.ItemSaver(
		fmt.Sprintf(":%d", config.ItemSaverPort))
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 1000,
		ItemChan:    itemChan}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}
