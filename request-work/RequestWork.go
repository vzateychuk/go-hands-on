package main

import (
	"log"
	"runtime"
)

func push(id string, from, to int, in <-chan bool, out chan<- int) {

	// Before sending anything, it waits for a request from the in channel
	for i := from; i <= to; i++ {
		<-in
		log.Println(id, i)
		out <- i
		runtime.Gosched()
	}
}

func main() {

	out := make(chan int, 100)
	in := make(chan bool, 100)
	go push("1", 1, 25, in, out)
	go push("2", 26, 50, in, out)
	go push("3", 51, 75, in, out)
	go push("4", 76, 100, in, out)

	sum := 0
	for i := 0; i < 100; i++ {
		in <- true
		pushed := <-out
		sum += pushed
	}
	close(in)
	close(out)
	log.Println("-------- SUM:", sum)
}
