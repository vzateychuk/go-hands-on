package myhttp

import (
	"encoding/json"
	"fmt"
	"http-db/db"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PeopleHandler struct {
	Name string
}

// Получить либо список данных либо одну запись и записать в response
func (handler *PeopleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	invokeCount.Inc()

	id := request.URL.Query().Get("id") // получить параметр GET запроса
	bytes := []byte{}
	var err error
	if strings.TrimSpace(id) == "" {
		people, err := db.GetPeople()
		if err == nil {
			bytes, err = json.MarshalIndent(people, "", "   ")
		}
	} else {
		uid, err := strconv.Atoi(id)
		if err != nil {
			log.Panic(err)
			fmt.Fprintln(writer, err)
		}
		person, err := db.GetPerson(uid)
		if err == nil {
			bytes, err = json.MarshalIndent(person, "", "   ")
		}
	}
	if err != nil {
		log.Panic(err)
		fmt.Fprintln(writer, err)
	} else {
		fmt.Fprintln(writer, string(bytes))
	}

	httpReqs.WithLabelValues("200", "GET").Inc()

	fmt.Fprintln(writer, "Handler:", handler.Name, "URL:", request.URL.String(), "id:", id)
}
