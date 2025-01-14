package main

import "fmt"

// 2. create concrete strategies

type Lfu struct {
}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by LFU strategy")
}
