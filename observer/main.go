package main

func main() {
	// Construct two DataListener observers and give each one a name
	lis1 := DataListener{Name: "Listener 1"}
	lis2 := DataListener{Name: "Listener 2"}

	// Create the DataSubject that the listeners will observe
	subj := &DataSubject{}

	// Register each listener with the DataSubject
	subj.registerObserver(lis1)
	subj.registerObserver(lis2)

	// Change the data in the DataSubject - this will cause the onUpdate method of each listener to be called
	subj.ChangeItem("Changed: Monday!")
	subj.ChangeItem("Changed: Tuesday!")

	// Unregister one of the observers and change the data
	subj.unregisterObserver(lis2)
	subj.ChangeItem("Changed: Wednesday!")
	subj.ChangeItem("Changed: Friday!")
}
