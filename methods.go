package main

import "fmt"

type Person struct {
	key	int
	name string
}

func (p *Person) setName(name string) {
	p.name = name
}

func main() {
	pers := Person{1, "Vlad"}
	pers.setName("VZateychuk")
	fmt.Println(pers)
}