package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func (cache *Cache) Add(key string, val []byte) error {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

	return nil
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}

	return entry.val, true
}

func NewCache(interval time.Duration) *Cache {
    cache := &Cache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}

    go cache.reapLoop(interval)

    return cache
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			cache.mu.Lock()
			for key, entry := range cache.entries {
				if time.Since(entry.createdAt) > interval {
					delete(cache.entries, key)
				}
			}
			cache.mu.Unlock()
		}
	}
}
