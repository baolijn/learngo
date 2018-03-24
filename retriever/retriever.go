package main

import (
	"learngo/retriever/mock"
	"fmt"
	"learngo/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(R Retriever) string  {
	return R.Get("https://www.baidu.com")
}

func main() {
	retriever := mock.Retriever{"eden is the best"}
	r := download(retriever)
	fmt.Println(r)

	realRetriever := real.Retriever{"libao"}
	fmt.Println(download(realRetriever))
}
