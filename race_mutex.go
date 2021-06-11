package main

import (
	"fmt"
	"sync"
)

func main() {
	counters := map[int]int{}
	mutex := sync.Mutex{} // создаем ссылку на mutex
	for i := 0; i < 5; i++ {
		go func(counters map[int]int, workerId int) {
			for j := 0; j < 5; j++ {
				mutex.Lock()              // Блокируем доступ к counters из разных goroutine
				counters[workerId*10+j]++ // читаем и записываем данные в map
				mutex.Unlock()            // Разблокируем доступ
			}
		}(counters, i)
	}

	fmt.Scanln()
	mutex.Lock()                             // Нужно и тут поскольку есть обращение к map-е counters, которое тоже может быть небезопасным
	fmt.Println("Counter result:", counters) // если не делать Lock, то race-condition detect обнаружит
	mutex.Unlock()                           // Разблокируем доступ из разных goroutine
}
