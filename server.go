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
	fmt.Println("i am a new branch called hotfix, i am here to merge to branch master")
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
