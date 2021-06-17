package main

import (
	"fmt"
)

// Демонстрация closure: функция возвращает функцию которую потом использует main
func decrementor(i int) func() int {
	var count = i
	return func() int {
		count--
		return count
	}
}

func main() {

	dec := decrementor(4)
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
	fmt.Println(dec())
}
