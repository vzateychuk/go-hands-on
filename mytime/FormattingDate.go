package main

import (
	"fmt"
	"time"
)

func timeToFormattedString() string {

	current := time.Now()
	formatted := fmt.Sprintf("%d:%d:%d %d/%d/%d",
		current.Hour(),
		current.Minute(),
		current.Second(),
		current.Year(),
		int(current.Month()),
		current.Day(),
	)
	fmt.Println(formatted)
	return formatted
}
