package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func worker1(ctx context.Context, workerNum int, out chan<- int) {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond // эмулируем случайную задержку
	fmt.Println(workerNum, "sleep", waitTime)
	select {
	case <-ctx.Done(): // здесь обрабатывается сигнал завершения (finish()). Выходим из функции
		return
	case <-time.After(waitTime): // в канал записывается результат работы по завершению задержки
		fmt.Println("worker1", workerNum, "done")
		out <- workerNum
	}
}

func main() {
	workTime := 50 * time.Millisecond
	ctx, _ := context.WithTimeout(context.Background(), workTime)
	result := make(chan int, 1)
	// Запускаем пакет work-в, у каждого случайная задержка
	for i := 0; i < 10; i++ {
		go worker1(ctx, i, result)
	}
	totalFound := 0
	// В цикле опрашиваем завершились ли worker1-ы до момента получения сигнала Done() контекста
LOOP:
	for {
		select {
		case <-ctx.Done():
			break LOOP
		case foundBy := <-result:
			totalFound++
			fmt.Println("result found by", foundBy)
		}
	}
	fmt.Println("totalFound:", totalFound)
	time.Sleep(time.Second)
}
