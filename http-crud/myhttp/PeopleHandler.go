package myhttp

import (
	"fmt"
	"net/http"
)

type PeopleHandler struct {
	Name string
}

// Структура PeopleHandler implements Handler поскольку для нее опреден метод ServerHTTP
func (handler *PeopleHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "PeopleName:", handler.Name, "URL:", request.URL.String()) // пишем каким обработчиком обработано
}
