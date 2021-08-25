package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// основные настройки к базе
	dsn := "user=postgres password=root host=127.0.0.1 port=5432 dbname=postgres sslmode=disable"

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalln("Connect error: ", err)
	}

	err = db.DB().Ping()
	if err != nil {
		log.Fatalln("Ping error: ", err)
	}

	handlers := &Handler{
		DB:   db,
		Tmpl: template.Must(template.ParseGlob("./html/*")),
	}

	// в целям упрощения примера пропущена авторизация и csrf
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
