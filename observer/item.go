package main

import "fmt"

type Product struct {
	observerList []Observer
	name         string
	inStock      bool
}

func newProduct(name string) *Product {
	return &Product{
		name: name,
	}
}

func (p *Product) updateStock() {
	fmt.Printf("The product %s is in stock.\n", p.name)
	p.inStock = true
	p.notifyAll()
}

func (p *Product) register(o Observer) {
	p.observerList = append(p.observerList, o)
}

func (p *Product) deregister(o Observer) {
	p.observerList = removeFromslice(p.observerList, o)
}

func (p *Product) notifyAll() {
	for _, observer := range p.observerList {
		observer.update(p.name)
	}
}

func removeFromslice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
