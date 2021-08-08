package main

import "fmt"

func main() {
	// IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback)

	// Use IterateBooks to iterate via anonymous function
	lib.IterateBooks(func(book Book) error {
		fmt.Println("Book author:", book.Author)
		return nil
	})

	// Traverse Books by BookIterator just created
	iter := lib.createIterator()
	for iter.hasNext() {
		book := iter.next()
		fmt.Printf("Book %+v\n", book)
	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
