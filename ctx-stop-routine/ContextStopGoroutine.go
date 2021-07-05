package main

import (
	"context"
	"log"
	"time"
)

func countNumbers(c context.Context, read chan int) {
	val := 0
	for {
		select {
			case <-c.Done():
				// In this select group, we have a case where we check whether the context is done, and if it is, 
				// we just break the loop and return the value we have counted so far
				read <- val
			default:
				// If the context is not done, we need to keep counting
				time.Sleep(time.Microsecond)
				val++			
		}
	}
}

func main() {
	read := make(chan int)
	// We need to be able to cancel the context, so we extend this simple context with a cancellable context:
	todo := context.TODO()
	ctx, stopFunc := context.WithCancel(todo)
	// call the counting routine
	go countNumbers(ctx, read)
	// At this point, we need a way to break the loop, so we will use the stop() function
	// returned by context.WithCancel, but we will do that inside another Goroutine. This
	// will stop the context after 300 milliseconds:
	go func() {
		time.Sleep(300 * time.Millisecond)
		stopFunc()
	}()

	// And at this point, we just need to wait for the message with the count to be received and log it
	val := <-read
	log.Println(val)
}