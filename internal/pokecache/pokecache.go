package pokecache

import (
	"time"
)

type Cache struct {
	mapCache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		mapCache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.mapCache[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Add(key string, data []byte) {
	c.mapCache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		data:      data,
	}
}

func (c *Cache) reapLoop (interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	conception := time.Now().UTC().Add(-interval)
	for key, entry := range c.mapCache {
		if entry.createdAt.Before(conception) {
			delete(c.mapCache, key)
		}
	}
}