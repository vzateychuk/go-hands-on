package main

import "fmt"

func swap(c *int, d *int) {
	fmt.Printf("before swap: c=%#v, *c=%#v, d=%#v, *d=%#v\n", c, *c, d, *d)
	// c,d = d,c // не работает
	*c, *d = *d, *c // вот так работает
	fmt.Printf("after swap: c=%#v, *c=%#v, d=%#v, *d=%#v\n", c, *c, d, *d)
}

func main() {
	a, b := 5, 10

	fmt.Printf("main before: a=%#v, &a=%#v, b=%#v, &b=%#v\n", a, &a, b, &b)
	swap(&a, &b) // call swap here
	fmt.Printf("main after: a=%#v, &a=%#v, b=%#v, &b=%#v\n", a, &a, b, &b)

	fmt.Println(a == 10, b == 5)
}
