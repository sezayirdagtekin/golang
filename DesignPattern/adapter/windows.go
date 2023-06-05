package main

import "fmt"

type windows struct {
	name    string
	version string
}

func (w *windows) insertInSquarePort() {
	fmt.Printf("Insert square port into windows machine name %s  version %s\n", w.name, w.version)
}
