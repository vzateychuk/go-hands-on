package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	iterAmt       = 7
	goroutinesAmt = 5
	quotaLimit    = 2
)

func startLimitedWorkerAsync(workerId int, wg *sync.WaitGroup, quotaCh chan struct{}) {
	quotaCh <- struct{}{} // goroutine захватывает свободный слот канала, блокируя выполнение других goroutine
	defer wg.Done()
	for j := 0; j < iterAmt; j++ {
		time.Sleep(time.Duration(rand.Intn(1000)+10) * time.Millisecond)
		fmt.Printf("workerId: %v, iter: %v\n", workerId, j)
		if j%2 == 0 {
			<-quotaCh             // освобождаем слот
			quotaCh <- struct{}{} // и тут же пробуем его захватить
		}
		runtime.Gosched() // передаем управление другим goroutine
	}
	<-quotaCh // освобождаем слот исполнения, занятый ранее, выбирая из нее структуру
}

func main() {
	wg := sync.WaitGroup{}
	quotaCh := make(chan struct{}, quotaLimit) // буф.канал с заданным ограничением
	for i := 0; i < goroutinesAmt; i++ {
		wg.Add(1)
		go startLimitedWorkerAsync(i, &wg, quotaCh) // передаем ограничивающий канал в асинх worker
	}
	time.Sleep(time.Millisecond)
	wg.Wait() // Ожидать завершения всех worker-в
}
