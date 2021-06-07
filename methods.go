package main

import "fmt"

type Person struct {
	key	int
	name string
}

type Account struct {
	Id int
	Number string
	Person
}

func (p *Person) setName(name string) {
	p.name = name
}

func main() {
	pers := new(Person)
	pers.setName("VZateychuk")

	var acc Account = Account{
		Id: 1,
		Number: "abc1234",
		Person: Person{
			key: 2,
			name: "",
		},
	}

	acc.setName("Ivan Ivanov")
	fmt.Println(acc)
}