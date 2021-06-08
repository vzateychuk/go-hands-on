package main

import (
	"fmt"
	"runtime"
	"strings"
)

const (
	ITERATIONS    = 7
	GOROUTNES_AMT = 5
)

func doSomeWork(in int) {

	for j := 0; j < ITERATIONS; j++ {
		fmt.Println(formatWork(in, j))
		// принудительно передаем управление другой goroutine
		runtime.Gosched()
	}
}

func formatWork(in int, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", GOROUTNES_AMT-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}

func main() {
	for i := 0; i < GOROUTNES_AMT; i++ {
		// приставка go инструкция выполнения goroutine, т.е. не блокировать главный поток выполнением doSomeWork
		go doSomeWork(i)
	}
	fmt.Scanln()
}
