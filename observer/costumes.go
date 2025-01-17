package main

import "fmt"

// 3. create concrete Observer

type Costumer struct {
	id string
}

func (c *Costumer) update(productName string) {
	fmt.Printf("Sending e-mail to costumer %s about product %s\n", c.id, productName)
}

func (c *Costumer) getID() string {
	return c.id
}
