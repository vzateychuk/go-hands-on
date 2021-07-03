package worker

import (
	"math/rand"
	"runtime"
	"time"
)

func DoWork(in <-chan int, out chan<- int) {
	sum := 0
	for i := range in {
		sum += i
	}
	time.Sleep( time.Duration(rand.Intn(100)+10) * time.Millisecond )
	runtime.Gosched()
	out <- sum
}

