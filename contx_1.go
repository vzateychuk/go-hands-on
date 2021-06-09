package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// объявляем ctx и finish - функция отмены. context.Background() - базовый contx
	ctx, finish := context.WithCancel(context.Background())
	result := make(chan int, 1) // канал передачи результата

	// запускаем worker - асинхронные запросы, передавая контекст, содержащий функцию отмены
	for i := 0; i < 10; i++ {
		go worker(ctx, i, result)
	}
	foundBy := <-result // дожидаемся первого результата
	fmt.Println("result found by: ", foundBy)
	finish() // функция отмены в контексте ctx. Обрабатывается в worker

	time.Sleep(time.Second) // ждем завершения главной goroutine
}

func worker(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond // эмулируем случайную задержку
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done(): // здесь обрабатывается сигнал завершения (finish()). Выходим из функции
		return
	case <-time.After(waitTime): // в канал записывается результат работы по завершению задержки
		fmt.Println("worker", workerNum, "done")
		out <- workerNum
	}
}
