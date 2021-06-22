package main

import (
	"flag"
	"fmt"
)

func main() {
	i := flag.Int("age", -1, "your age")
	n := flag.String("name", "", "your first name")
	b := flag.Bool("married", false, "are you married?")
	flag.Parse()
	fmt.Println("Name: ", *n)
	fmt.Println("Age: ", *i)
	fmt.Println("Married: ", *b)
}
