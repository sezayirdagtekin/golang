package main

import (
	"fmt"
	"sync"
)

type Mylogger struct {
	loglevel int
}

func (l *Mylogger) log(msg string) {
	fmt.Println(l.loglevel, ":", msg)
}

func (l *Mylogger) SetLogLevel(level int) {
	l.loglevel = level
}

var logger *Mylogger

var onceTimes sync.Once

func getLoggerInstance() *Mylogger {

	onceTimes.Do(func() {
		fmt.Println("Creating logger instance")
		logger = &Mylogger{}
	})

	fmt.Println("Returning logger instance")
	return logger
}
