package main

import (
	"net/http"

	"vez/depinv/log"
	"vez/depinv/store"
)

func main() {

	logger := LoggerAdapter(log.PrintMessage)
	store := store.NewSimpleDataStore()
	logic := NewSimpleLogic(logger, store)
	contr := NewController(logger, logic)
	http.HandleFunc("/hello", contr.SayHello)
	http.ListenAndServe(":8080", nil)
}
