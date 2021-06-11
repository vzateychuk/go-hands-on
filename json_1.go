package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id       int    // public, начинается с заглавной буквы
	Username string // public
	phone    string // private, начинается с маленькой буквы
}

var jsonString = `{ "id": 42, "username": "vzateychuk", "phone": "1234" }`

func main() {
	data := []byte(jsonString)

	user := User{}
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("unmarshal user: ", user) // поле phone не будет демаршализировано, поскольку private

	user.phone = "345678"
	bytes, err := json.Marshal(user) // поле phone не будет маршализировано поскольку private
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("marshal bytes:", string(bytes))
}
