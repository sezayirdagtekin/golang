package main

func main() {

	w := &windows{
		name:    "Dell windows",
		version: "home",
	}
	mac := &mac{
		name:  "MacBook",
		model: "Air M2",
	}

	c := &client{}
	winAdap := &windowsAdapter{
		windowsMachine: w,
	}

	c.insertSquareUsbInComputer(mac)

	c.insertSquareUsbInComputer(winAdap)

}
