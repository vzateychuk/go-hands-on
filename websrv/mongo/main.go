package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"html/template"
	"net/http"
)

func main() {
	// подключение к Mongo
	sess, err := mgo.Dial("mongodb://root:example@localhost:27017")
	__err_panic(err)

	// если коллекции не будет, то она создасться автоматически
	collection := sess.DB("coursera").C("items")

	// для Mongo вставляем демо-запись если коллекция пуста
	if n, _ := collection.Count(); n == 0 {
		err1 := collection.Insert(NewItem("mongodb", "Рассказать про монгу", ""))
		__err_panic(err1)
		err2 := collection.Insert(NewItem("redis", "Рассказать про redis", "vez"))
		__err_panic(err2)
	}

	handlers := &Handler{
		Items: collection,
		Tmpl:  template.Must(template.ParseGlob("./templates/*")),
	}

	// Объявляем http роутеры, связывая Path и обработчики
	//в целям упрощения примера пропущена авторизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.List).Methods("GET")
	r.HandleFunc("/items", handlers.List).Methods("GET")
	r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
	r.HandleFunc("/items/new", handlers.Add).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", r)
}
