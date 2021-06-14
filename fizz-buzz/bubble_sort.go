package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	numberAmount = 1_000
)

func main() {
	// создать и заполнить массив случайными числами
	nums := make([]int, numberAmount)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(nums); i++ {
		rnd := rand.Intn(numberAmount)
		nums[i] = rnd
		fmt.Printf("%d,\n", rnd)
	}
	fmt.Println("=== Data prepared ===")
	// пока остаются переставленные значения
	for swapped := true; swapped; {
		swapped = false
		// проходим в цикле по парам значений и переставляем
		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i-1], nums[i] = nums[i], nums[i-1]
				swapped = true
			}
		}
	}
	fmt.Println("Result:", nums)
}
