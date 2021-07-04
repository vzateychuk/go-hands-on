package main

import (
	"fmt"
	"sync"
)

func readThem(ch chan int, wg *sync.WaitGroup) {
	// Будет вычитывать из канала до его закрытия
	for i := range ch {
		fmt.Println(i)
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go readThem(ch, wg)
	fmt.Println("Put 1")
	ch <- 1
 	fmt.Println("Put 2")
	ch <- 2
	fmt.Println("Put 3")
	ch <- 3
	fmt.Println("Put 4")
	ch <- 4
	fmt.Println("Put 5")
	ch <- 5
    // close(ch) // Если мы не закроем канал, то заблокирется обе goroutine (как main, так и readThem), поскольку wg.Wait()
	wg.Wait()
	fmt.Println("Reach the end")
}