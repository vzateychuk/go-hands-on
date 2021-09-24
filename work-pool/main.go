package main

import (
	"fmt"
	"math/rand"
	"time"
	"work-pool/work"
)

const goroutineAmount = 4

func main() {
	worker := work.GetWorkerInstance(goroutineAmount)

	for i := 0; i < 5; i++ {
		var taskAmt = rand.Intn(20) + 1
		workItems := work.NewWorkItems(taskAmt)
		fmt.Printf("====> Main: Start batch: [%d], tasks: [%d]\n", i, taskAmt)

		for _, wrk := range workItems {
			worker.DoWork(wrk)
		}

		fmt.Printf("<==== Main: Finish batch: [%d]\n", i)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(30 * time.Second)
}
