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
    // adds new entry to the cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
    // gets an entry from the cache
    // return true if exists, false otherwise
    return []byte{}, false
}

func (c *Cache) reapLoop() {
    // remove entries older than 'interval'
}

func NewCache(interval time.Duration) *Cache {
    return &Cache{}
}
