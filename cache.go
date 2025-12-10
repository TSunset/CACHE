package cache

import (
	"errors"
	"sync"
	"time"
)

type item struct {
	value interface{}
	exAt  time.Time
}

type Cache struct {
	cache map[string]item
	mu    sync.RWMutex
}

func New() *Cache {
	return &Cache{
		cache: make(map[string]item),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = item{
		value: value,
		exAt:  time.Now().Add(ttl),
	}

}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.RLock()
	val, ok := c.cache[key]
	c.mu.RUnlock()

	if !ok {
		return nil, errors.New("not found")
	}

	if time.Now().After(val.exAt) {
		c.mu.Lock()
		delete(c.cache, key)
		c.mu.Unlock()

		return nil, errors.New("time exit")
	}

	return val.value, nil
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.cache, key)
}
