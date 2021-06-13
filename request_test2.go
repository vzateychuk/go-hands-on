package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Card struct {
	PaymentApiURL string // внешний сервис
}
type CheckoutResult struct {
	Status  int
	Balance int
	Err     string
}

// Функция структуры Сard, выполняющая оплату и возвращающая CheckoutResult
func (c *Card) Checkout(id string) (*CheckoutResult, error) {
	url := c.PaymentApiURL + "?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	// вычитываем Body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &CheckoutResult{}
	// распаковываем JSON
	err = json.Unmarshal(data, result)
	return result, err
}

// Функция возвращающая ответ мокового сервера
func MockServiceHandler(w http.ResponseWriter, r *http.Request) {
	key := r.FormValue("id")
	switch key {
	case "42":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":200, "balance": "100500"}`)
	case "100500":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":400, "balance": "bad_balance"}`)
	case "broken_json":
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, `{"status":400}`)
	case "internal_error":
		fallthrough
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}

type TestCase1 struct {
	ID      string
	Result  CheckoutResult // ожидаемый &CheckoutResult
	IsError bool
}

func TestCardCheckout(t *testing.T) {
	testCases := []TestCase1{
		TestCase1{ // успешный тест кейс
			ID: "42",
			Result: CheckoutResult{
				Status:  400,
				Balance: 0,
				Err:     "bad_balance",
			},
			IsError: false,
		},
		TestCase1{
			ID: "100500",
			Result: CheckoutResult{
				Status:  400,
				Balance: 0,
				Err:     "bad_balance",
			},
			IsError: false,
		},
		TestCase1{
			ID:      "broken_json",
			Result:  CheckoutResult{},
			IsError: true,
		},
		TestCase1{
			ID:      "internal_error",
			Result:  CheckoutResult{},
			IsError: true,
		},
	}

	// Поднимает mock сервис на случайном порту, передавая MockServiceHandler, который будет обрабатывать все запросы.
	testServer := httptest.NewServer(http.HandlerFunc(MockServiceHandler))

	// итеририруемся по тесткейсам, отправляя запросы в mock сервис
	for caseNum, testCase := range testCases {
		card := &Card{
			PaymentApiURL: testServer.URL, // Передаем адрес нашего mock сервера
		}
		result, err := card.Checkout(testCase.ID)
		// assert что функция не возвращает err когда не должна
		if err != nil && !testCase.IsError {
			t.Errorf("[%d] unexpected error: %#v", caseNum, err)
		}
		// assert что функция возвращает err когда должна
		if err == nil && testCase.IsError {
			t.Errorf("[%d expected error, got nil", caseNum)
		}
		// assert response
		if !reflect.DeepEqual(testCase.Result, result) {
			t.Errorf("[%d] wrong result, expected: %#v, got: %#v", caseNum, testCase.Result, result)
		}
	}
	testServer.Close()
}
