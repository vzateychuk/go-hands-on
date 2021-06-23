package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigsChan := make(chan os.Signal, 1)
	doneChan := make(chan bool)
	signal.Notify(sigsChan, syscall.SIGINT)
	go func() {
		for {
			s := <-sigsChan
			fmt.Println("Received signal: ", s)
			switch s {
			case syscall.SIGINT:
				fmt.Println("My process has been interrupted. Someone might of pressed CTRLC")
				fmt.Println("***Some clean up is occuring***")
				doneChan <- true
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	<-doneChan
	fmt.Println("Out of here")
}
