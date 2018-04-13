package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
	"regexp"
	"learngo/crawler/scheduler"
	"learngo/crawler/persist"
)

var (profileRe  = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	cityUrlRe = regexp.MustCompile(
		`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func main() {
	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	//e := engine.ConcurrentEngine{
	//	Scheduler: &scheduler.SimpleScheduler{},
	//	WorkerCount: 1}
	itemChan, err := persist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 2000,
		ItemChan:    itemChan}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	//e.Run(engine.Request{
	//	Url:	"http://www.zhenai.com/zhenghun/shanghai",
	//	ParserFunc: parser.ParseCity,
	//})
	//e.Run(engine.Request{
	//	Url:        "http://album.zhenai.com/u/109732029",
	//	ParserFunc: parser.ProfileParser(""),
	//})
}
