package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Handler struct {
	Sess  *mgo.Session
	Items *mgo.Collection
	Tmpl  *template.Template
}

func (handler *Handler) List(w http.ResponseWriter, r *http.Request) {

	items := []*Item{}

	// bson.M{} - это типа условия для поиска
	err := handler.Items.Find(bson.M{}).All(&items)
	__err_panic(err)

	err = handler.Tmpl.ExecuteTemplate(w, "index.html", struct{ Items []*Item }{Items: items})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := handler.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) Add(w http.ResponseWriter, r *http.Request) {

	newItem := NewItem(r.FormValue("title"), r.FormValue("description"), "")
	/*
		bson.M{
			"_id":         bson.NewObjectId(),
			"title":       r.FormValue("title"),
			"description": r.FormValue("description"),
			"some_filed":  123,
		}
	*/
	err := handler.Items.Insert(newItem)
	__err_panic(err)

	fmt.Println("Insert - LastInsertId:", newItem.Id)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (handler *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "bad id", 500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])

	post := &Item{}
	err := handler.Items.Find(bson.M{"_id": id}).One(&post)

	err = handler.Tmpl.ExecuteTemplate(w, "edit.html", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handler *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "bad id", 500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])

	post := &Item{}
	err := handler.Items.Find(bson.M{"_id": id}).One(&post)

	post.Title = r.FormValue("title")
	post.Description = r.FormValue("description")
	post.Updated = "vez"

	// Можно провести Update передавая структуру post:
	// err = handler.Items.Update(bson.M{"_id": id}, &post)
	// Либо можно выполнить Update с помощью bson
	err = handler.Items.Update(
		bson.M{"_id": id},
		bson.M{
			"title":       r.FormValue("title"),
			"description": r.FormValue("description"),
			"updated":     "vez",
			"newField":    123,
		})
	affected := 1
	if err == mgo.ErrNotFound {
		affected = 0
	} else if err != nil {
		__err_panic(err)
	}

	fmt.Println("Update - RowsAffected", affected)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (handler *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if !bson.IsObjectIdHex(vars["id"]) {
		http.Error(w, "bad id", 500)
		return
	}
	id := bson.ObjectIdHex(vars["id"])

	err := handler.Items.Remove(bson.M{"_id": id})
	affected := 1
	if err == mgo.ErrNotFound {
		affected = 0
	} else if err != nil {
		__err_panic(err)
	}

	w.Header().Set("Content-type", "application/json")
	resp := `{"affected": ` + strconv.Itoa(int(affected)) + `}`
	w.Write([]byte(resp))
}

// Только для демонстрационных целей.
// В prod, ошибка должна всегда явно обрабатываться
func __err_panic(err error) {
	if err != nil {
		panic(err)
	}
}
