package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func main() {

	http.HandleFunc("/", GetUser)

}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	key := request.FormValue("id")
	if key == "42" {
		writer.WriteHeader(http.StatusOK)
		io.WriteString(writer, `{"status":200, "resp": {"user": 42}}`)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		io.WriteString(writer, `{"status":500, "err": "some_error"}`)
	}
}

type TestCase struct {
	ID         string // идентификатор юзера для формирования запроса (FormValue("id"))
	Response   string // ожидаемый ResponseBody ответа
	StatusCode int    // ожидаемый statusCode
}

func TestGetUser(t *testing.T) {
	// Подготовавливаем тест кейсы
	testCases := []TestCase{
		TestCase{ // успешный тест кейс
			ID:         "42",
			Response:   `{"status":200, "resp": {"user": 42}}`,
			StatusCode: http.StatusOK,
		},
		TestCase{ // тест кейс с ответом 500 сервера
			ID:         "500",
			Response:   `{"status":500, "err": "some_error"}`,
			StatusCode: http.StatusInternalServerError,
		},
	}
	// Итерируемся по ТестКейсам
	for caseNum, testCase := range testCases {
		// arrange
		url := "http://apiuser?id=" + testCase.ID
		req := httptest.NewRequest("GET", url, nil)
		recorder := httptest.NewRecorder() // Записывает все что делает тестируемая функция с writer

		// act
		GetUser(recorder, req)

		// assert StatusCode
		if recorder.Code != testCase.StatusCode {
			t.Errorf("[%d] wrong StatusCode: expected %d, got %d", caseNum, testCase.StatusCode, recorder.Code)
		}
		// assert body
		resp := recorder.Result()
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(body)
		if bodyStr != testCase.Response {
			t.Errorf("[%d] wrong Response: expected %+v, got %+v", caseNum, testCase.Response, bodyStr)
		}
	}
}
