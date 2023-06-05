package main

func main() {
	// Use the Facade Cafe API to create coffee drinks
	// instead of directly interacting with the complex Coffee API

	// Make an 8 ounce Americano
	makeAmericano(8.0)

	// Make a 12 ounce Latte
	makeLatte(12.0, true)

	/*  OUTPUT
	PS C:\development\DesignPattern\facade> go run .\main.go  .\CoffeeMachine.go .\coffee.go
	Making an Americano
	--------------------
	Starting coffee order with beans: 5 and grind level 2
	Grinding the beans: 5 beans at 2 granularity
	Adding hot water: 8
	Ending coffee order
	Americano is ready!

	Making a Latte
	--------------------
	Starting coffee order with beans: 7.5 and grind level 4
	Grinding the beans: 7.5 beans at 4 granularity
	Adding hot water: 12
	Adding milk: 3 oz
	Foam setting: true
	Ending coffee order
	Latte is ready!
	PS C:\development\DesignPattern\facade>

	*/
}
