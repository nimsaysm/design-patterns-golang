package main

// 1. create instance that will publish events

type Subject interface {
	register(observer Observer)
	deregister(observer Observer)
	notifyAll()
}
