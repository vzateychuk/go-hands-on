package main

import (
	"log"
	"vez/splitwork/worker"
)

/* number of workers and the range for the numbers to process */
func sum(workers, from, to int) int {
	// create channels
	in := make(chan int, 4)
	out := make(chan int, 4)
	// create workers
	for i := 0; i < workers; i++ {
		go worker.DoWork(in, out)
	}
	// create a loop to send all the numbers to the in channel
	for i := from; i < to; i++ {
		in <- i
	}
	// As we sent all the numbers, we now need to receive the partial sums back, but before that we need to notify the function that the numbers to sum are finished
	close(in)
	// And then perform the sum of the partials
	sum := 0
	for i:=0;i<workers; i++ {
		sum += <-out
	}
	// finally, close the out channel and return the result
	close(out)
	return sum
}

func main() {
	res := sum(2,1,100)
	log.Println("Result:", res)
}