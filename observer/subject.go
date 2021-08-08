package main

import "fmt"

// The interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unregisterObserver(obs observer)
	notifyAll()
}

// The DataSubject have a list of listeners and a field that gets changed, triggering them
type DataSubject struct {
	observers []DataListener
	field     string
}

// The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) ChangeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// This function adds an observer to the list
func (ds *DataSubject) registerObserver(dl DataListener) {
	ds.observers = append(ds.observers, dl)
}

// This function removes an observer from the list
func (ds *DataSubject) unregisterObserver(dl DataListener) {
	var newObservers []DataListener
	for _, obs := range ds.observers {
		if obs.Name != dl.Name {
			newObservers = append(newObservers, obs)
		}
	}
	ds.observers = newObservers
	fmt.Println("Listener: ", dl.Name, "unregistered!")
}

// The notifyAll function calls all the listeners with the changed data (ds.field)
func (ds *DataSubject) notifyAll() {
	for _, dataListener := range ds.observers {
		dataListener.onUpdate(ds.field)
	}
}
