package pokeapi

import (
	"sync"
	"time"
)

type cache struct {
	entry    map[string]cacheEntry
	interval time.Duration
	sync.RWMutex
}
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *cache {
	c := cache{
		entry:    map[string]cacheEntry{},
		interval: interval,
	}
	go c.reapLoop()
	return &c
}

func (c *cache) Add(key string, val []byte) {
	defer c.Unlock()
	c.Lock()
	c.entry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}
func (c *cache) Get(key string) (val []byte, found bool) {
	defer c.Unlock()
	c.Lock()
	en, found := c.entry[key]
	if !found {
		return val, found
	}
	return en.val, found
}

func (c *cache) reapLoop() {
	t := time.Tick(c.interval)
	for range t {
		c.Lock()
		for key, val := range c.entry {
			if time.Since(val.createdAt) > c.interval {
				delete(c.entry, key)
			}
		}
		c.Unlock()
	}

}
