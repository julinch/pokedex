package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries map[string]CacheEntry
	mu      sync.Mutex
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	var cache Cache
	cache.Entries = make(map[string]CacheEntry)
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Entries[key] = CacheEntry{CreatedAt: time.Now(), Val: val}
}

func (c *Cache) Get(key string) (val []byte, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var entry CacheEntry
	entry, ok = c.Entries[key]
	if !ok {
		return nil, false
	}
	val = entry.Val
	return val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.Entries {
			if time.Since(value.CreatedAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}
