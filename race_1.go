package main

import "fmt"

func main() {
	counters := map[int]int{}
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, workerId int) {
			for j := 0; j < 5; j++ {
				counters[workerId*10+j]++ // читаем и записываем данные в map
			}
		}(counters, i)
	}

	fmt.Scanln() // чтобы успели goroutine стартовать
	fmt.Println("Counter result:", counters)
}
