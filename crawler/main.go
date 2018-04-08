package main

import (
	"learngo/crawler/engine"
	"learngo/crawler/zhenai/parser"
)

func main() {
	engine.SimpleEngine{}.Run(engine.Request{
		Url:	"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}

