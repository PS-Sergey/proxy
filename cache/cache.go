package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheSize int
	items     map[string]Item
	mutex     sync.Mutex
}

type Item struct {
	createTime time.Time
	data       []byte
}

func NewCache(cacheSize int) *Cache {
	return &Cache{
		cacheSize: cacheSize,
		items:     make(map[string]Item),
	}
}

func (cache *Cache) GetValue(key string) ([]byte, bool) {
	item, ok := cache.items[key]
	if !ok {
		return nil, ok
	}
	return item.data, ok
}

func (cache *Cache) SetValue(key string, value []byte) {
	if cache.cacheSize <= 0 {
		return
	}
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if len(cache.items) == cache.cacheSize {
		oldestTime := time.Now()
		var oldestKey string
		for key, item := range cache.items {
			if item.createTime.Before(oldestTime) {
				oldestTime = item.createTime
				oldestKey = key
			}
		}
		delete(cache.items, oldestKey)
	}
	cache.items[key] = Item{createTime: time.Now(), data: value}
}
