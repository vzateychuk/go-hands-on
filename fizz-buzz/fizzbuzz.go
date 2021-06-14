package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(i int) string {
	switch {
	case i%3 == 0 && i%5 == 0:
		return "FizzBuzz"
	case i%3 == 0:
		return "Fizz"
	case i%3 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(i)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
