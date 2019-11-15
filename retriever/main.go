package main

import (
	"fmt"
	"go-exercise/retriever/mock"
	real2 "go-exercise/retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://niracler.com")
}

func main() {
	var r1, r2 Retriever
	r1 = mock.Retriever{"this is a fack"}
	fmt.Println(download(r1))
	r2 = real2.Retriever{}
	fmt.Println(download(r2))
}
