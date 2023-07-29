package main

import (
	"errors"
)

func NewSimpleLogic(log MyLogger, ds MyDataStore) *SimpleLogic {

	return &SimpleLogic{
		log:   log,
		store: ds,
	}
}

type SimpleLogic struct {
	log   MyLogger
	store MyDataStore
}

func (sl *SimpleLogic) SayHello(userId string) (string, error) {
	sl.log.MyLog("invoked SayHello for: " + userId)

	if name, ok := sl.store.UserNameForID(userId); ok {
		sl.log.MyLog("found user: " + name)
		return "Hello, " + name, nil
	}

	return "", errors.New("user not found")
}

func (sl *SimpleLogic) SayGoodBye(userId string) (string, error) {
	sl.log.MyLog("invoked SayGoodBye for: " + userId)

	if name, ok := sl.store.UserNameForID(userId); ok {
		sl.log.MyLog("found user: " + name)
		return "Good bye, " + name, nil
	}

	return "", errors.New("user not found")
}
