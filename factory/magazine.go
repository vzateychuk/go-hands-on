package main

import "fmt"

// A magazine struct and embed the publication interface
type magazine struct {
	publication
}

// A Stringer interface gives a string representation of the type
func (m magazine) String() string {
	return fmt.Sprintf("This is a magazine named %s", m.name)
}

// the createMagazine function returns a new Magazine object
func createMagazine(name string, pages int, publisher string) iPublication {
	return &magazine{
		publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		}}
}
