package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	mapCache map[string]cacheEntry
	mapMutex *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func NewCache() Cache {
	return Cache{
		mapCache: make(map[string]cacheEntry),
		mapMutex: &sync.Mutex{},
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mapMutex.Lock()
	defer c.mapMutex.Unlock()
	entry, ok := c.mapCache[key]
	if !ok {
		return nil, false
	}
	return entry.data, true
}

func (c *Cache) Add(key string, data []byte) {
	c.mapMutex.Lock()
	defer c.mapMutex.Unlock()
	c.mapCache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		data:      data,
	}
}