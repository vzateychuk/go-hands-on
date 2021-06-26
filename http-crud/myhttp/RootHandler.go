package myhttp

import (
	"fmt"
	"net/http"
)

type RootHandler struct {
	Name string
}

// Структура PeopleHandler implements Handler поскольку для нее опреден метод ServerHTTP
func (handler *RootHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "RootName:", handler.Name, "URL:", request.URL.String()) // пишем каким обработчиком обработано
}
