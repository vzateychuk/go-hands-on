package main

import "gopkg.in/mgo.v2/bson"

type Item struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Updated     string        `json:"updated" bson:"updated"`
}

func NewItem(title, desc, updated string) *Item {

	return &Item{
		Id:          bson.NewObjectId(),
		Title:       title,
		Description: desc,
		Updated:     updated,
	}
}
