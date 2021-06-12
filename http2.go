package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
	Name string
}

func (handler *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Name:", handler.Name, "URL:", request.URL.String())
}

func main() {
	testHandler := &MyHandler{Name: "testHandler"} // будет обрабатывать запросы на test
	http.Handle("/test/", testHandler)

	rootHandler := &MyHandler{Name: "rootHandler"}
	http.Handle("/", rootHandler)

	fmt.Println("Server started :8080")
	http.ListenAndServe(":8080", nil)
}
