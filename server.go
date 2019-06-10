package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, request.Header)
}

func main() {
	fmt.Println("i am on branch iss01")
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.3:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
