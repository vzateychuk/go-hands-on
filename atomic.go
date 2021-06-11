package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var total int64

func inc() {
	atomic.AddInt64(&total, 1) // важно что обращение к переменной идет по ссылке
}
func main() {
	for i := 0; i < 1000; i++ {
		go inc()
	}
	// Важно: если не дать время завершиться всем goroutine, результат будет случайным!
	time.Sleep(10 * time.Microsecond)
	fmt.Println(total)
}
