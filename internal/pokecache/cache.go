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

func NewCache(interval time.Duration) *Cache {
    newCache := &Cache{}
    newCache.cache = map[string]cacheEntry{}

    go newCache.reapLoop(interval)

    return newCache
}

func (c *Cache) Add(key string, value []byte) {
    c.Lock()
    defer c.Unlock()
    c.cache[key] = cacheEntry{
	createdAt: time.Now(),
	val: value,
    }
}

func (c *Cache) Get(key string) ([]byte, bool) {
    c.Lock()
    defer c.Unlock()
    value, exists := c.cache[key]
    return value.val, exists
}

func (c *Cache) reapLoop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for range ticker.C {
	c.reap(interval)
    }
}

func (c *Cache) reap(interval time.Duration) {
    c.Lock()
    defer c.Unlock()
    for k, v := range c.cache { 
	elapsed := time.Since(v.createdAt)
	if elapsed > interval {
	    //fmt.Printf("elapsed time is %v which is greater than interval %v, deleting cache entry %s with creation time %v\n", elapsed, interval, k, v.createdAt)
	    delete(c.cache, k)
	}
    }
}

