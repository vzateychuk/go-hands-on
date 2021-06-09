package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	i := 0
	for tickerTime := range ticker.C {
		i++
		fmt.Println("step: ", i, ", time: ", tickerTime)
		// останавливаем если прошло более 5 тиков
		if i >= 5 {
			ticker.Stop()
			break // важно выйти из цикла, иначе будет deadlock
		}
	}
	fmt.Println("total: ", i)

	// Тикер time.Tick(time.Second) не может быть остановлен и будет работать до явного завершения программы
	// Можно использовать например как запуск агента мониторинга
	c := time.Tick(time.Second)
	for tickTime := range c {
		fmt.Println("step: ", i, ", time: ", tickTime)
		break // не остановит тикер time.Tick, он будет работать даже при выходе из цикла
	}
}
