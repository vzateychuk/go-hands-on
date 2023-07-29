package main

import "net/http"

type Logic interface {
	SayHello(userId string) (string, error)
}

type Controller struct {
	log   MyLogger
	logic Logic
}

func NewController(log MyLogger, logic Logic) Controller {
	return Controller{
		log:   log,
		logic: logic,
	}
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.log.MyLog("In Controller.SayHello")

	userId := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}
