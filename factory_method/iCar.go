package main

// 1. create the interface with all methods that an object must have

type ICar interface {
	setName(name string)
	getName() string
}
