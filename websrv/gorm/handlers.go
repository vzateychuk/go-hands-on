package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Item struct {
	Id          int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Title       string
	Description string
	Updated     string `sql:"null"`
}

// Здесь можно задать любое имя для таблицы
func (i *Item) TableName() string {
	return "items"
}

func (i *Item) BeforeSave() (err error) {
	fmt.Println("trigger on before save")
	return
}

type Handler struct {
	DB   *gorm.DB
	Tmpl *template.Template
}

func (handl *Handler) List(w http.ResponseWriter, r *http.Request) {

	items := []*Item{}

	db := handl.DB.Find(&items)
	err := db.Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = handl.Tmpl.ExecuteTemplate(w,
		"index.html",
		struct{ Items []*Item }{Items: items})
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (handl *Handler) AddForm(w http.ResponseWriter, r *http.Request) {
	err := handl.Tmpl.ExecuteTemplate(w, "create.html", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Edit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := &Item{}

	db := h.DB.Find(post, id)
	err = db.Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("Record not found", id)
	} else {
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	err = h.Tmpl.ExecuteTemplate(w, "edit.html", post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post := &Item{}
	h.DB.Find(post, id)

	post.Title = r.FormValue("title")
	post.Description = r.FormValue("description")
	post.Updated = "updater-name"

	db := h.DB.Save(post)
	err = db.Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	affected := db.RowsAffected

	fmt.Println("Update - RowsAffected", affected)

	http.Redirect(w, r, "/", http.StatusFound)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	db := h.DB.Delete(&Item{Id: id})
	err = db.Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	affected := db.RowsAffected

	fmt.Println("Delete - RowsAffected", affected)

	w.Header().Set("Content-type", "application/json")
	resp := `{"affected": ` + strconv.Itoa(int(affected)) + `}`
	w.Write([]byte(resp))
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {

	newItem := &Item{
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
	}
	db := h.DB.Create(&newItem)
	err := db.Error
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	affected := db.RowsAffected

	fmt.Println("Insert - RowsAffected", affected, "LastInsertId: ", newItem.Id)

	http.Redirect(w, r, "/", http.StatusFound)
}
