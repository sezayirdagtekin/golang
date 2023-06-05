package main

import "fmt"

type windowsAdapter struct {
	windowsMachine *windows
}

func (w *windowsAdapter) insertInSquarePort() {

	fmt.Println("Calling windowsAdapter...")
	w.windowsMachine.insertInSquarePort()
}
