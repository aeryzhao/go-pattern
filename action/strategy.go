package main

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type FIFO struct{}

func (f *FIFO) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}

type LRU struct{}

func (l *LRU) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

type LFU struct{}

func (l *LFU) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

type Cache struct {
	storage      map[string]string
	capacity     int
	maxCapacity  int
	evictionAlgo EvictionAlgo
}

func NewCache(maxCapacity int, evictionAlgo EvictionAlgo) *Cache {
	return &Cache{
		storage:      make(map[string]string),
		capacity:     0,
		maxCapacity:  maxCapacity,
		evictionAlgo: evictionAlgo,
	}
}

func (c *Cache) SetEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	fmt.Println(len(c.storage), c.capacity)
	if c.capacity >= c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

func (c *Cache) del(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

func main() {
	lru := &LRU{}
	cache := NewCache(2, lru)
	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")
	f := &FIFO{}
	cache.SetEvictionAlgo(f)
	cache.add("d", "4")
	lfu := &LFU{}
	cache.SetEvictionAlgo(lfu)
	cache.add("e", "5")
}
