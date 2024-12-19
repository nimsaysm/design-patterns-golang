package main

import "fmt"

// 5. create client code

func main() {
	bmw, _ := getCar("Bmw")
	ferrari, _ := getCar("Ferrari")

	printInfo(bmw)
	printInfo(ferrari)
}

func printInfo(c ICar) {
	fmt.Println("Car: ", c.getName())
}
