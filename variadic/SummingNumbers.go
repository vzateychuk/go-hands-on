package main

import (
	"fmt"
)

// sum function accepts a variadic argument of type int
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	i := []int{5, 10, 15}
	fmt.Println(sum(i...))
	fmt.Println(sum(5, 4))
}
