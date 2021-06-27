package myhttp

import (
	"html/template"
	"http-db/model"
	"net/http"
	"strconv"
)

type RootHandler struct {
	tpl *template.Template
}

// Конструктор rootHandler по заданному пути к html шаблонам
func NewRootHandler(tplPath string) (*RootHandler, error) {
	tmpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return nil, err
	}
	return &RootHandler{tmpl}, nil
}

// Здесь в случае метода Post получаются данные из формы и возвращаются в виде Person
func (handler *RootHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	person := model.Person{}
	if r.Method == http.MethodPost {
		err := r.ParseForm() // parse sent form
		if err != nil {
			w.WriteHeader(400)
			return
		}
		// If the form gets parsed correctly, we can proceed
		person.Firstname = r.Form.Get("firstname")
		person.Lastname = r.Form.Get("lastname")
		person.Age, _ = strconv.Atoi(r.Form.Get("age"))
	}
	handler.tpl.Execute(w, person) // we can finally return the page
}
