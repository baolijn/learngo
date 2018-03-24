package main

import (
	"learngo/retriever/mock"
	"fmt"
	"learngo/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(R Retriever) string  {
	return R.Get("https://www.baidu.com")
}

func inspect(r Retriever)  {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println(v.Contents)
	case real.Retriever:
		fmt.Println(v.UserAgent, v.TimeOut)
	}
}
func main() {
	var r Retriever
	r = mock.Retriever{"eden is the best"}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r,r)
	inspect(r)
	r = real.Retriever{"libao", time.Minute}
	fmt.Printf("%T %v\n", r,r)

	//type assertion
	mockRetriever, ok := r.(mock.Retriever)
	if ok{
		fmt.Println(mockRetriever)
	}else{
		panic("err")
	}

	//fmt.Println(download(realRetriever))
}
