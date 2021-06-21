package main

import (
	"fmt"
	"time"
)

func main() {

	dur := 6*time.Hour + 6*time.Minute + 6*time.Second

	start := time.Now()
	finish := start.Add(dur)
	fmt.Println("+6 hours from now will be at: ", finish.Format(time.ANSIC))
}
