package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "FizzBuzz"
	} else if i%3 == 0 {
		return "Fizz"
	} else if i%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(i)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
