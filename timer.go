package main

import (
	"fmt"
	"time"
)

// функция эмулирует выполнение длительной операции (2 сек)
func longSQLQuery(ch chan bool) {
	time.Sleep(2 * time.Second)
	ch <- true
}

// функция возвращает bool канал в котором true по завершении длительной операции
func chanQuery() chan bool {
	ch := make(chan bool, 1)
	go longSQLQuery(ch)
	return ch
}

func main() {
	timer := time.NewTimer(1 * time.Second) // задаем величину таймаута 1 сек
	select {
	case <-timer.C:
		fmt.Println("timer.C timeout")
	case <-time.After(time.Minute):
		// Таймаут через минуту, но проблема: будет работать пока не сработает After,
		//даже после остановки программы и не соберется сборщиком мусора
		fmt.Println("After Minute timeout")
	case result := <-chanQuery():
		// по сигналу завершения долгой операции, освобождаем ресурсы таймера
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("operation result", result)
	}
}
