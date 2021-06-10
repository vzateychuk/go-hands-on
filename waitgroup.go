package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

const (
	iterAmount   = 7
	goroutineAmt = 5
)

func startWorkerMy(workerId int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for i := 0; i < iterAmount; i++ {
		time.Sleep(time.Duration(rand.Intn(1000)+10) * time.Millisecond)
		fmt.Printf("workerId: %v, iter: %v\n", workerId, i)
		runtime.Gosched()
	}
}

func main() {
	waitGroup := sync.WaitGroup{} // Счетчик goroutine, которые должны будут завершиться одновременно
	for i := 0; i < goroutineAmt; i++ {
		/*
			waitGroup.Add можно было бы сделать внутри 'startWorkerMy', но тогда необходимо приостановить goroutine 'main'
			(с помощью time.Sleep ниже) иначе выполнение waitGroup.Wait() в 'main' может произойти раньше чем будет выполнено
			добавление waitGroup.Add внутри 'startWorkerMy' и тогда 'main' завершит выполнение раньше
		*/
		waitGroup.Add(1)
		go startWorkerMy(i, &waitGroup)
	}
	// time.Sleep(time.Millisecond)

	waitGroup.Wait() // Ожидаем
}
