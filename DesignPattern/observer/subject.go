package main

// Define the interface for the observable type
type observable interface {
	registerObserver(obs observer)
	unRegisterObserver(obs observer)
	notifyAll()
}

type DataSubject struct {
	observers []DataListener
	field     string
}

// The ChangeItem function will cause the Listeners to be called
func (ds *DataSubject) changeItem(data string) {
	ds.field = data
	ds.notifyAll()
}

// This function adds an observer to the list
func (ds *DataSubject) registerObserver(o DataListener) {
	ds.observers = append(ds.observers, o)
}

// This function removes an observer from the list
func (ds *DataSubject) uRegisterObserver(o DataListener) {
	var newObs []DataListener
	for _, obs := range ds.observers {
		if o.Name != obs.Name {
			newObs = append(newObs, obs)
		}
	}
	ds.observers = newObs
}

// The notifyAll function calls all the listeners with the changed data
func (ds *DataSubject) notifyAll() {

	for _, ob := range ds.observers {
		ob.onUpdate(ds.field)
	}
}
