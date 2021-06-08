package main

import "fmt"

func main() {
	dataCh := make(chan int)   // Канал передачи данных
	cancelCh := make(chan int) // Канал для выхода из функции
	// функция генератор значений. Остановить ее можно передачей значения в канале отмены cancelCh
	go func(cancel chan int, data chan int) {
		val := 0
		for {
			select {
			case data <- val:
				val++ // наращиваем значение в цикле и передаем в канал
			case <-cancel:
				return // если в канале 'cancel' обнаружится чтото, выходим из функциии
			}
		}
	}(cancelCh, dataCh)
	// в цикле зачитываем значения канала dataCh пока значение не будет больше трех
	for curVal := range dataCh {
		fmt.Println("read: ", curVal)
		if curVal > 45 { // Когда значение больше трех, отправляем значение в канал отмены
			fmt.Println("send cancel")
			cancelCh <- curVal
			break
		}
	}
}
