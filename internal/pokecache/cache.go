package pokecache

import (
    "sync"
    "time"
)


type cacheEntry struct {
    createdAt	time.Time
    val		[]byte
}

type Cache struct { 
    sync.Mutex
    cache map[string]cacheEntry
}

func (c *Cache) Add(key string, val []byte) {
    c.Lock()
    newEntry := cacheEntry{
	createdAt: time.Now(),
	val: val,
    }
    c.cache[key] = newEntry
    c.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
    _, exists := c.cache[key]
    if exists {
	return c.cache[key].val, true
    }

    return []byte{}, false
}

func (c *Cache) reapLoop() {
    // remove entries older than 'interval'
}

func NewCache(interval time.Duration) *Cache {
    newCache := &Cache{}
    newCache.cache = map[string]cacheEntry{}
    return newCache
}
