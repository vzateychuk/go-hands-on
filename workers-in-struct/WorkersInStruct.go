package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


type Worker struct {
	inChan, outChan chan int
	subwrk     	int	// number of subworkers
	mtx     	*sync.Mutex
}

// Note that the function is meant to be used directly and not as a Goroutine, as it itself creates a new Goroutine.
func (wrk *Worker) createReadAndSumGoroutine() {
	// increment amount of subworkers
	wrk.subwrk++ 
	// create anonymous goroutine which is calculating sum 
	go func() {
		partial := 0
		// считаем сумму из канала in
		for i := range wrk.inChan {
			partial += i
			time.Sleep(time.Duration(rand.Intn(10)+1) * time.Millisecond)	// эмулируем "длительную" операцию
		}
		wrk.outChan <- partial

		// we've locked the routine, reduced the counter on the sub-workers safely, 
		// and then, in case all the workers have terminated, we've closed the output channel
		wrk.mtx.Lock()
			wrk.subwrk--	// уменьшаем количество subroutine
			if wrk.subwrk == 0 {
				close(wrk.outChan)
			}
		wrk.mtx.Unlock()
	}()
}

// a function that's able to return the sum
func (wrk *Worker) gatherResult() int {

	// we add 1 to it as we will spawn only one routine
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Считаем total, складывая все что придет из channel out	
	// we have looped until the out channel is closed by one of the sub-workers (in readThem()).
	total := 0
	go func() {
		for i:= range wrk.outChan{
			total += i
		}
		wg.Done()
	}()
	// Ждем пока закончит goroutine и возвращаем результат
	wg.Wait()
	return total
}

func main() {
	mtx := &sync.Mutex{}
	in := make(chan int, 100)
	workersNum := 10
	out := make(chan int)
	wrk := Worker{inChan: in, outChan:out, mtx:mtx}

	// Now create a loop where you call the createReadAndSumGoroutine() method workersNum times. 
	// This will create some sub-workers which will process data from in channel:
	for i:=1; i<=workersNum; i++ {
		wrk.createReadAndSumGoroutine()
	}

	// Now send the numbers to be processed to the 'in' channel
	for i:=1; i<=100_000; i++ {
		in <- i
	}

	// Close the channel to notify that all the numbers have been sent
	close(in)
	
	// Wait for the result and print it out
	res := wrk.gatherResult()
	fmt.Println(res)
}