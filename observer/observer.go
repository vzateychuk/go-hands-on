package main

import "fmt"

// The interface for an observer type
type observer interface {
	onUpdate(data string)
}

// DataListener has name which identifies the dataListener
type DataListener struct {
	Name string
}

// an onUpdate function to implements observer
func (d *DataListener) onUpdate(data string) {
	fmt.Println("Listener: ", d.Name, "got data change:", data)
}
