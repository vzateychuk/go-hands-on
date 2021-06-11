package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

func main() {
	mutex := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go incCounter(mutex)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("finish:", counter)
}

func incCounter(mutex sync.Mutex) {
	mutex.Lock()
	counter++
	mutex.Unlock()
}
