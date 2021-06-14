package main

import "fmt"

func swap(a *int, b *int) {
	fmt.Printf("before swap: a=%#v, *a=%#v, b=%#v, *b=%#v\n", a, *a, b, *b)
	// a,b = b,a // не работает
	*a, *b = *b, *a // вот так работает
	fmt.Printf("after swap: a=%#v, *a=%#v, b=%#v, *b=%#v\n", a, *a, b, *b)
}

func main() {
	a, b := 5, 10

	fmt.Printf("main before: a=%#v, &a=%#v, b=%#v, &b=%#v\n", a, &a, b, &b)
	swap(&a, &b) // call swap here
	fmt.Printf("main after: a=%#v, &a=%#v, b=%#v, &b=%#v\n", a, &a, b, &b)

	fmt.Println(a == 10, b == 5)
}
