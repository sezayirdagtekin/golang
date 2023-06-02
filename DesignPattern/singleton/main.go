package main

import "fmt"

func main() {
	/*
		log := getLoggerInstance()
		log.loglevel = 2
		log.log("hi there")

		log.loglevel = 1
		log.log("hello")

		log.loglevel = 3
		log.log("how are you?")

	*/

	//create several goroutines that try to get the
	for i := 0; i < 5; i++ {
		go getLoggerInstance()
	}
	fmt.Scanln()

	/* output
	PS C:\development\DesignPattern\singleton> go run .\main.go .\MyLogger.go
	Creating logger instance
	Returning logger instance
	Returning logger instance
	Returning logger instance
	Returning logger instance
	Returning logger instance

	*/

}
