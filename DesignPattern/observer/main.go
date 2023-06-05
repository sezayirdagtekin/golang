package main

func main() {

	// Construct two DataListener observers and
	// give each one a name
	listener1 := DataListener{
		Name: "Listener 1",
	}

	listener2 := DataListener{
		Name: "Listener 2",
	}

	// Create the DataSubject that the listeners will observe
	subj := &DataSubject{}
	subj.registerObserver(listener1)
	subj.registerObserver(listener2)

	// Change the data in the DataSubject - this will cause the
	// onUpdate method of each listener to be called
	subj.changeItem("Monday")
	subj.changeItem("Wednesday")
	subj.changeItem("Friday")

	// Try to unregister one of the observers
	subj.uRegisterObserver(listener2)
	subj.changeItem("Saturday")

	/* OUTPUT:
	PS C:\development\DesignPattern\observer> go run .\main.go  .\observer.go .\subject.go
	Listener: Listener 1 got data change: Monday
	Listener: Listener 2 got data change: Monday
	Listener: Listener 1 got data change: Wednesday
	Listener: Listener 2 got data change: Wednesday
	Listener: Listener 1 got data change: Friday
	Listener: Listener 2 got data change: Friday
	Listener: Listener 1 got data change: Saturday


	*/
}
