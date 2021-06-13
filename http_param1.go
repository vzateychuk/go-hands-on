package main

import (
	"fmt"
	"net/http"
)

func myhanler(writer http.ResponseWriter, request *http.Request) {
	param := request.URL.Query().Get("param") // получить параметр из запроса
	key := request.FormValue("key")           // получим используя функцию FormValue
	fmt.Fprintln(writer, "param:", param, "key:", key)
}

func main() {
	http.HandleFunc("/", myhanler)

	fmt.Println("Server started")
	http.ListenAndServe(":8080", nil)
}
