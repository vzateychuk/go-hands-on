package main

import (
	"fmt"
	"time"
)

func main() {
	current := time.Now()
	// format to: "02:49:21 21/06/2021."
	formatted := timeToFormattedString(current)
	fmt.Println(formatted)
}

func timeToFormattedString(mytime time.Time) string {

	return fmt.Sprintf("%d:%d:%d %d/%d/%d",
		mytime.Hour(),
		mytime.Minute(),
		mytime.Second(),
		mytime.Year(),
		int(mytime.Month()),
		mytime.Day(),
	)
}
