package main

import (
	"fmt"
	"log"
	"net/http"
)

type indexHandler struct {
}

func (i indexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "hello world")
	log.Panicln("http 8080")
	if err != nil {
		return
	}
}

func main() {
	fmt.Print("GoLang后台开始启动")
	mux := http.NewServeMux()
	mux.Handle("/", &indexHandler{})

	e := http.ListenAndServe(":8080", mux)
	if e != nil {
		fmt.Print(e)
	}
}
