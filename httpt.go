package main

import (
	"log"
	"net/http"
)

type server int

// ServeHTTP(ResponseWriter, *Request) 实现这个接口的s才能作为2th参
func (h *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	contentLength, err := w.Write([]byte("hello world"))
	if err != nil {
		panic(err)
	}
	log.Println("content Length:", contentLength) //11 about
}
func main() {
	var s server
	http.ListenAndServe("localhost:9999", &s)
}
