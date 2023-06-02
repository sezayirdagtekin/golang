package main

import "fmt"

func main() {

	mag1, _ := newPublication("magazine", "Tyme", 50, "The Tymes")
	mag2, _ := newPublication("magazine", "Lyfe", 40, "Lyfe Inc")
	news1, _ := newPublication("newspaper", "The Herald", 60, "Heralders")
	news2, _ := newPublication("newspaper", "The Standard", 30, "Standarders")

	printPublication(mag1)
	printPublication(mag2)
	printPublication(news1)
	printPublication(news2)
}

func printPublication(pub iPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.getName())
	fmt.Printf("Pages: %d\n", pub.getPages())
	fmt.Printf("Publisher: %s\n", pub.getPublisher())
	fmt.Printf("--------------------\n")

}
