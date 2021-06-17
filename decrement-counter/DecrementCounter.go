package main

import (
	"fmt"
	"time"
)

// Демонстрация closure: функция возвращает функцию которую потом использует main
func decrementor(i int) func() int {
	var count = i
	return func() int {
		count -= 1
		return count
	}
}

func main() {

	dec := decrementor(4)
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	time.Sleep(time.Second)
}
