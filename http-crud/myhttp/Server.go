package myhttp

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RunServer(port int) {

	recordUptime()
	prometheus.MustRegister(httpReqs)

	http.Handle("/metrics", promhttp.Handler())

	testHandler := &PeopleHandler{Name: "peopleHandler"} // обработчик
	http.Handle("/people/", testHandler)

	rootHandler, err := NewRootHandler("./html/index.html")
	if err != nil {
		log.Panic(err)
	}
	http.Handle("/", rootHandler) // все запросы кроме /people/ будут обрабатываться этим обработчиком
	// по запросу /form возвращается форма из файла
	http.HandleFunc("/form",
		func(writer http.ResponseWriter, request *http.Request) {
			http.ServeFile(writer, request, "./html/form.html")
		})
	log.Printf("Server started :%d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil) // стартуем сервер

}
