package main

import "fmt"

// Define a newspaper type and embed the publication interface
type newspaper struct {
	publication
}

// Define a stringNewsPaper interface that gives a string representation of the type
func (n newspaper) stringNewsPaper() string {
	return fmt.Sprintf("This is a newspaper named %s", n.name)
}

// the createNewspaper function returns a new Newspaper object
func createNewsPaper(name string, pages int, publisher string) iPublication {

	return &newspaper{
		publication: publication{
			name:      name,
			pages:     pages,
			publisher: publisher,
		},
	}

}
