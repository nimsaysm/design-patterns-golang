package main

// 2. Observer must subscribe to the subject events to be notified when they happen

type Observer interface {
	update(string)
	getID() string
}
