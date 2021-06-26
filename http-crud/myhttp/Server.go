package myhttp

import (
	"fmt"
	"log"
	"net/http"
)

func RunServer(port int) {
	testHandler := &PeopleHandler{Name: "peopleHandler"} // обработчик
	http.Handle("/people/", testHandler)

	rootHandler := &RootHandler{Name: "rootHandler"}
	http.Handle("/", rootHandler) // все запросы кроме /people/ будут обрабатываться этим обработчиком

	log.Printf("Server started :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // стартуем сервер

}
