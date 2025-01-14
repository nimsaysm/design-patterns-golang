package main

// 4. create client code

func main() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("x", "1")
	cache.add("y", "2")
	cache.add("z", "3")

	// will change algorithm because the cache is in your max capacity
	lru := &Lru{}
	cache.setEviction(lru)

	cache.add("a", "4")

	fifo := &Fifo{}
	cache.setEviction(fifo)

	cache.add("b", "5")
}
