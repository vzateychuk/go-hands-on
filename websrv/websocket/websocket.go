package main

import (
	"encoding/json"
	"fmt"
	"github.com/icrowley/fake"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	// Просто выводит шаблон в index.html
	tmpl := template.Must(template.ParseFiles("index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	})

	// Здесь мы пробуем upgrade соединения до web-socket
	http.HandleFunc("/notifications", func(w http.ResponseWriter, r *http.Request) {
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		go sendNewMsgNotifications(ws)
	})
	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func sendNewMsgNotifications(client *websocket.Conn) {
	// Каждые 3 секунды записываем данные в веб-сокет
	ticker := time.NewTicker(3 * time.Second)

	for i := 0; ; i++ {
		// Получаем writer для записи текстовых данных в сокет
		w, err := client.NextWriter(websocket.TextMessage)
		// если не удалось получить Writer, останавливаем ticker и выходим
		if err != nil {
			ticker.Stop()
			break
		}
		// создаем новый message и записываем в Writer
		msg := newMessage(i)
		w.Write(msg)
		w.Close()
		// Ожидаем когда сработает ticker
		<-ticker.C
	}
}

func newMessage(i int) []byte {
	data, _ := json.Marshal(map[string]string{
		"email":   fake.EmailAddress(),
		"name":    strconv.Itoa(i) + " " + fake.FirstName() + " " + fake.LastName(),
		"subject": fake.Product() + " " + fake.Model(),
	})
	return data
}
