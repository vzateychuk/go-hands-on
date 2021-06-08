package main

import "fmt"

func main() {
	ch1 := make(chan int)

	// запуск goroutine "на канале" ch1
	go func(in chan int) {
		// инициируем val из канала `in`
		val := <-in
		fmt.Println("GO: get from chan: ", val)
		fmt.Println("GO: after read from chan")
	}(ch1)

	// кладем значение 42 в канал ch1
	ch1 <- 10
	fmt.Println("Main: after put to chan")

	fmt.Scanln()
}
