package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

const goroutineAmount = 3

func doWork(workerId int, input string) string {
	waitTime := time.Duration(rand.Intn(100)+10) * time.Millisecond // эмулируем случайную задержку
	time.Sleep(waitTime)
	return fmt.Sprintln(strings.Repeat("  ", workerId), "█",
		strings.Repeat("  ", goroutineAmount-workerId),
		"workerId", workerId, "finished", input)
}

func startWorker(workerId int, in <-chan string) {
	for input := range in { // пока в канале есть значения вычитываем
		fmt.Println(doWork(workerId, input))
		runtime.Gosched() // передаем выполнение следующей goroutine
	}
	fmt.Printf("=== worker: %v stop\n", workerId) // сообщает о завершении worker
}

func main() {
	runtime.GOMAXPROCS(1)               // количество cpu которые могут быть использованы
	workerInput := make(chan string, 2) // канал которым будут пользоваться goroutine
	// стартует pool worker-в зачитывающих значения из канала
	for i := 0; i < goroutineAmount; i++ {
		go startWorker(i, workerInput)
	}
	// заполняем канал значениями из массива months
	months := []string{
		"Январь", "Февраль", "Март", "Апрель", "Май", "Июнь", "Июль", "Август", "Сентябрь", "Октябрь", "Ноябрь", "Декабрь",
	}
	for _, month := range months {
		workerInput <- month
	}
	// Если не закрыть канал, то main завершится раньше, и может привести к "зависнувшим" workers
	close(workerInput) // закрываем канал (когда будут выбраны все значения).
	time.Sleep(time.Second)
}
