package main

import (
	"fmt"
	"time"
)

func main() {
	Current := time.Now()
	Los_Angeles, _ := time.LoadLocation("America/Los_Angeles")
	fmt.Println("The local current time is:", Current.Format(time.ANSIC))
	fmt.Println("The time in Los_Angeles is: ", Current.In(Los_Angeles).Format(time.ANSIC))
}
