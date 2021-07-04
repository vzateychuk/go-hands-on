package main

import "log"

func readThem(in, done chan string) {
	for s := range in {
		log.Println(s)
	}
	done <- "done"
}

func main() {
	// create the necessary channels and use them
	in, done := make(chan string), make(chan string)
	go readThem(in, done)
	// create a set of strings and loop over them, sending each string to the channel
	strs := []string{"a","b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s
	}
	// close the channel you used to send the messages and wait for the done signal
	close(in)
	<-done
}