package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func post(poster Poster) {
	poster.Post("http://www.imooc.com", map[string]string{
		"name":   "ccmouse",
		"course": "golang",
	})
}
type RetrieverPoster interface {
	Retriever
	Poster
}
const url = "http://www.imooc.com"
func session(s RetrieverPoster) string {
	s.Post(url,map[string]string{
		"contents":"another faked imooc.com",
	})
    return s.Get(url)
}
func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	fmt.Printf("%T %v\n", r, r, )
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	fmt.Printf("%T %v\n", r, r, )
	//fmt.Println(download(r))
}
