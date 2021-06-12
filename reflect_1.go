package main

import (
	"fmt"
	"reflect"
)

type MyUser struct {
	ID       int
	RealName string
	Login    string
	Flags    int
}

func PrintReflect(u interface{}) error {
	object := reflect.ValueOf(u).Elem()

	fmt.Printf("%T have %d fields:\n", u, object.NumField())
	for i := 0; i < object.NumField(); i++ {
		fieldValue := object.Field(i)
		fieldType := object.Type().Field(i)

		fmt.Printf("\tname: %v, type: %v, value: %v, tag: `%v`\n",
			fieldType.Name,
			fieldType.Type.Kind(),
			fieldValue,
			fieldType.Tag)
	}
	return nil
}

func main() {
	user := MyUser{
		ID:       42,
		RealName: "vzateychuk",
		Login:    "vzateychuk",
		Flags:    32,
	}
	err := PrintReflect(&user)
	if err != nil {
		panic(err)
	}
}
