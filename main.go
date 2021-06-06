package main

import (
	"fmt"
)

const GLOBAL_LIMIT = 100
const MAX_CACHE_SIZE = 10 * GLOBAL_LIMIT

func main() {
	// 
	var users map[string]string = map[string]string{
		"name": "Vlad",
		"lastname": "Zateychuk",
	}
	fmt.Printf("%d %+v\n", len(users), users)

	// создание с нужной емкостью
	profiles := make(map[string]string, 10)
	fmt.Printf("%d %+v\n", len(profiles), profiles)
	// если ключ не удалось найти, вернётся значение по умолчанию для типа
	mName := profiles["unknown"]
	fmt.Printf("mName: %v\n", mName)

	// проверка на существование ключа
	mName, isNameExists := profiles["somename"]
	fmt.Println(isNameExists)

}
