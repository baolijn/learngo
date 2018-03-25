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

type Poster interface {
	Post(url string,
		forms map[string]string) string
}

type RetriverPoster interface {
	Retriever
	Poster
}

const url  = "https://www.baidu.com"

func download(R Retriever) string  {
	return R.Get(url)
}

func post(p Poster)  {
	p.Post(url, map[string]string{
		"name": "libao",
		"course": "golang",
	})
}

func session(r RetriverPoster) string  {
	r.Post(url, map[string]string{
		"contents":"another fake",
	})
	return r.Get(url)
}
func inspect(r Retriever)  {
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println(v.Contents)
	case real.Retriever:
		fmt.Println(v.UserAgent, v.TimeOut)
	}
}
func main() {
	var r Retriever
	retriever := mock.Retriever{"this is a fake"}
	//fmt.Println(download(r))
	//fmt.Printf("%T %v\n", r,r)
	inspect(r)
	r = real.Retriever{"libao", time.Minute}
	fmt.Printf("%T %v\n", r,r)

	//type assertion
	//mockRetriever, ok := r.(mock.Retriever)
	//if ok{
	//	fmt.Println(mockRetriever)
	//}else{
	//	panic("err")
	//}
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
	//fmt.Println(download(realRetriever))
}
