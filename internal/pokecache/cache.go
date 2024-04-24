package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mux      *sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{entries: map[string]cacheEntry{}, interval: interval}
	c.mux = &sync.RWMutex{}
	go func() {
		c.reapLoop()
	}()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.entries[key]
	if ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		<-ticker.C
		{
			c.mux.RLock()
			for key, v := range c.entries {
				t := time.Since(v.createdAt)
				if t > c.interval {
					delete(c.entries, key)
				}

			}
			c.mux.RUnlock()
		}
	}

}
