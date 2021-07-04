/*
Generates files with random integer values (0..100)
*/
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
)

func createFile(fileName string, nums []int) {
	f, err := os.Create(fileName)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	for i := 0; i < len(nums); i++ {	
		_, err = f.WriteString(fmt.Sprintf("%d\n", nums[i])) // writing...
		if err != nil {
			log.Printf("Error writing string: %v", err)
		}
	}
}

func makeIntArray(seed int64) []int {
	rand.Seed(seed)
	buf := make([]int, 100)
	for i := 0; i < 100; i++ {
		buf[i] = rand.Intn(100)
	}
	return buf
}

func main() {
	buf1 := makeIntArray(42)
	createFile("input1.dat", buf1)

	buf2 := makeIntArray(24)
	createFile("input2.dat", buf2)
}
