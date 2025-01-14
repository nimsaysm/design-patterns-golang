package main

// 1. create interface to be followed by main class
// the interface can change the algorithm in execution time

type Eviction interface {
	evict(c *Cache)
}
