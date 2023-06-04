package main

import (
	"fmt"
	"net/http"
	"time"
)

func doThingThatShouldBeLimited() string {
	waitSec := 10
	time.Sleep(time.Duration(waitSec) * time.Second)
	return fmt.Sprintf("done waiting: %d sec", waitSec)
}

func main() {
	bp := NewBackPressure(2)

	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		err := bp.Process(func() {
			w.Write([]byte(doThingThatShouldBeLimited()))
		})

		if err != nil {
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("Too many requests"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
