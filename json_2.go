package main

import (
	"encoding/json"
	"fmt"
)

// сразу настраивается сеарилизация/десереализация
type UserDto struct {
	Id       int    `json:"user_id,string"` // в json это поле будет называться user_id (string)
	Username string `json:","`              // ','-имя поля не будет переопределяться
	Address  string `json:",omitempty"`     // 'omitempty'- при сеарилизации пропускать, если пустое значение
	Company  string `json:"-"`              // знак "-" означает что поле вообще не сеарилизуется/десереализуется
}

func main() {
	user := UserDto{
		Id:       24,
		Username: "vlad",
		Address:  "",
		Company:  "VEZ",
	}
	bytes, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
