package main

import (
	"encoding/json"
	"fmt"
)

// JSON-ы с произвольной структурой
var jsonStr = `[
	{"uid": 14, "username": "vovk", "phone": "1234"},
	{"sid": "abc42", "address": "none", "company": "VES" }
]`

func main() {
	data := []byte(jsonStr)

	var user interface{} // Переменная типа "пустой интерфейс"
	err := json.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	// десереализуется в масив Map-в произвольной структуры
	fmt.Printf("unmarshal in empty interface: \n%#v\n", user)

	var address = map[string]interface{}{
		"sid":     "ef24",
		"address": "Moscow",
		"company": "EVA",
	}
	bytes, err := json.Marshal(address) // сериализуется в валидный map
	fmt.Printf("marshal to JSON:\n%#v\n", string(bytes))
}
