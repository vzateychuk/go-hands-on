package main

import (
	"fmt"
	"net/http"
	"time"
)

func runServer(port string) {
	mux := http.NewServeMux() // Создаем и настраиваем мультиплексор
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Port=", port, "URL=", request.URL.String())
	})

	server := http.Server{ // Создаем и настраиваем Http listener
		Addr:         port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server at", port)
	err := server.ListenAndServe() // стартуем сервер
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Стартуем сервера в отдельных goroutin-ах
	go runServer(":8080") // для порта 8080
	go runServer(":8181") // для порта 8181
	fmt.Scanln()          // не даем программе завершиться и завершить goroutine
}
