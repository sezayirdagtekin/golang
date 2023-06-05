package main

import "fmt"

// Iterate using a callback function
func main() {

	//1. way:  use IterateBooks to iterate via a callback function
	lib.IterateBooks(myBookCallback) //lib defined as  global library in books.go

	//2. way:  Use IterateBooks to iterate via anonymous function

	lib.IterateBooks(func(book Book) error {
		fmt.Println("Book author:", book.Author)
		return nil
	})

	//3. way: Create a BookIterator
	it := lib.createIterator()
	for it.hasNext() {
		book := it.next()
		fmt.Printf("Book %+v\n :", book)

	}
}

// This callback function processes an individual Book object
func myBookCallback(b Book) error {
	fmt.Println("Book title:", b.Name)
	return nil
}
