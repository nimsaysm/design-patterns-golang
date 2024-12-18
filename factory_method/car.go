package main

// 2. create concrete product - struct that implements iCar interface

type Car struct {
	name string
}

func (m *Car) setName(name string) {
	m.name = name
}

func (m *Car) getName() string {
	return m.name
}
