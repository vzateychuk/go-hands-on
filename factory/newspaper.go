package main

import "fmt"

// A newspaper type and embed the publication interface
type newspaper struct {
	publication
}

// A Stringer interface that gives a string representation of the type
func (n newspaper) String() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

// the createNewspaper function returns a new Newspaper object
func createNewspaper(name string, pages int, publisher string) iPublication {

	return &newspaper{
		publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		}}
}
