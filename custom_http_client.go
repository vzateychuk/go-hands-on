package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

/*
curl -v -X GET -H "Content-Type: application/json" -d '{ "id":2, "user":"vzateychuk" }' http://localhost:8080
*/

func runGetDefault() {
	url := "http://127.0.0.1:8080/?param1=123&param2=test"
	resp, err := http.Get(url) // получили http.Response
	defer resp.Body.Close()    // Важно всегда нужно закрывать запрос
	if err != nil {
		fmt.Println(err)
		return
	}
	respBody, err := ioutil.ReadAll(resp.Body) // считываем response.Body в байтовый массив
	fmt.Println("runGetDefault(): ", string(respBody))
}

func runGetFullReq() {
	request := &http.Request{
		Method: http.MethodGet,
		Header: http.Header{
			"User-Agent": {"vzateychuk/golang"},
		},
	}
	// Распарсить url чтобы можно было бы добавить параметров запроса
	request.URL, _ = url.Parse("http://127.0.0.1:8080/?id=42")
	request.URL.Query().Set("user", "vez") // добавляем параметры запроса ("&user=vez")
	// Используем default Http клиент для HTTP запроса и получения response
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close() // не забываем закрыть Body

	body, err := ioutil.ReadAll(response.Body)
	fmt.Printf("getFullReq: response: \n%#v\n", string(body))
}

/*
Использование default http-client нерекомендуется из за невозможности задать timeout.
Т.е. если сервер поломается, default http-client просто "зависнет".
тобы этого избежать используем настраиваемый http клиент: http.Transport
*/
func runTransportAndPost() {
	// структура Transport
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       60 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: time.Second,
	}
	// Создаем клиента с тайамаутами и ссылкой на transport
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: transport,
	}
	// создаем реквест с данными в body, headers и т.п.
	data := `{"id":42, "user":"newuser"}`
	body := bytes.NewBufferString(data) // io.Readerf
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1:8080/raw_body", body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(data)))
	// Выполняем http запрос к помощью созданного и настроенного клиента
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close() // не забываем закрывать body
	respBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("runTransportClient, %#v\n", string(respBody))
}

func startServer() {
	// Обработка GET на '/'
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "GET-handler. Incoming:\n%#v\n", request)
		fmt.Fprintf(writer, "GET-handler. Url:\n%#v\n", request.URL)
	})
	// Обработка POST на '/raw_body'
	http.HandleFunc("/raw_body", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body) // считали body в байтовый массив
		defer r.Body.Close()                // важно всегда закрывать поток Body
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Fprintf(w, "POST-handler. Raw body: %s\n", string(body))
	})

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	go startServer()
	fmt.Scanln()
	runGetDefault()
	fmt.Scanln()
	runGetFullReq()
	fmt.Scanln()
	runTransportAndPost()
	fmt.Scanln()
	fmt.Println("Finished")
}
