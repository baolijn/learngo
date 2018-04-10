package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"learngo/crawler/scheduler"
	"learngo/crawler/persist"
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount: 100}
	e := engine.ConcurrentEngine{
			Scheduler: &scheduler.QueuedScheduler{},
			WorkerCount: 1000,
			ItemChan:persist.ItemSaver()}
	e.Run(engine.Request{
		Url:	"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
}

