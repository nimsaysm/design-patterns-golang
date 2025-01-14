package main

import "fmt"

// 2. create concrete strategies

type Lru struct {
}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by LRU strategy")
}
