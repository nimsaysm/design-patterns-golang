package main

import "fmt"

// 2. create concrete strategies

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO strategy")
}
