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
    //fmt.Printf("cache entry %s created at %v\n", key, newEntry.createdAt)
    c.cache[key] = newEntry
    c.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.Lock()
    _, exists := c.cache[key]
    c.Unlock()
    if exists {
	return c.cache[key].val, true
    }

    return []byte{}, false
}

func (c *Cache) reapLoop() {
    limit := 5 * time.Second

    c.Lock()
    for k, v := range c.cache { 
	elapsed := time.Since(v.createdAt)
	if elapsed > limit {
	    //fmt.Printf("elapsed time is %v which is greater than limit %v, deleting cache entry %s with creation time %v\n", elapsed, limit, k, v.createdAt)
	    delete(c.cache, k)
	}
    }
    c.Unlock()
}

func NewCache(interval time.Duration) *Cache {
    newCache := &Cache{}
    newCache.cache = map[string]cacheEntry{}
    //newCache.reapLoop()
    ticker := time.NewTicker(interval)
    go func() { 
	for range ticker.C {
	    newCache.reapLoop()
	}
    }()
    return newCache
}
