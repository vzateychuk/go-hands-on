package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)

	ch1 <- 1 // в канале ch1 есть значение

	select {
	case val := <-ch1:
		// сработает когда в канале ch1 есть значения
		fmt.Println("ch1 val: ", val) // 1
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		// сработает когда ни в одном канале нет значений, если default упустить, заблокируется(то при пустых каналах)
		fmt.Println("default case")
	}
}
