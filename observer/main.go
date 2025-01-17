package main

// 4. create client code

func main() {
	mouse := newProduct("Mouse")

	observerTyler := &Costumer{id: "tyler@gmail.com"}
	observerBrian := &Costumer{id: "brian@gmail.com"}

	mouse.register(observerTyler)
	mouse.register(observerBrian)

	mouse.updateStock()
}
