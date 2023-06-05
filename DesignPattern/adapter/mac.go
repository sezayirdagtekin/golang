package main

import "fmt"

type mac struct {
	name  string
	model string
}

func (m *mac) insertInSquarePort() {
	fmt.Printf("Insert square port into mac machine name %s %s\n", m.name, m.model)
}
