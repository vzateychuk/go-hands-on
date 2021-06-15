package main

import "fmt"

func main() {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	result := append(week[6:], week[0:6]...)
	fmt.Println(result)
}
