package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	Date := time.Now()
	s := strconv.Itoa(Date.Year()) + "_" + Date.Month().String() + "_" + strconv.Itoa(Date.Day())
	fmt.Println(s)
}
