package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world!") // можно вывод сделать с помощью fmt.Fprint
	writer.Write([]byte("!!!"))        // либо можно использовать writer и записывать массив байт
}

func main() {
	http.HandleFunc("/page", // Обработка одной страницы '/page'
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "Single page:", request.URL.String())
		},
	)

	http.HandleFunc("/pages/", // Обработка любых запросов начинающихсяс '/pages/'
		func(writer http.ResponseWriter, request *http.Request) {
			fmt.Fprintln(writer, "Multiple pages:", request.URL.String())
		})

	http.HandleFunc("/", handler) // задается фунция для обработки запросов на '/'
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
