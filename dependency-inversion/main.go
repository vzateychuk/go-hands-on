package main

import (
	"net/http"

	"vez/depinv/log"
	"vez/depinv/store"
)

func main() {

	logger := LoggerAdapter(log.PrintMessage)
	ds := store.NewSimpleDataStore()
	dataStore := DataStoreAdapter(ds.GetUserById)
	logic := NewSimpleLogic(logger, dataStore)
	contr := NewController(logger, logic)
	http.HandleFunc("/hello", contr.SayHello)
	http.ListenAndServe(":8080", nil)
}
