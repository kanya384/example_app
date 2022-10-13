package memcache

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var (
	ErrNotFound = errors.New("not found")
)

type Cache struct {
	sync.RWMutex
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
	items             map[string]Item
	done              chan struct{}
}

type Item struct {
	Value      interface{}
	CreatedAt  time.Time
	Expiration int64
}

func New(defaultExpiration, cleanupInterval time.Duration) *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
		done:              make(chan struct{}),
	}

	if cleanupInterval > 0 {
		cache.startGC()
	}

	return &cache
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}

	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		CreatedAt:  time.Now(),
	}
}

func (c *Cache) Get(key string) (value interface{}, err error) {
	c.RLock()
	defer c.RUnlock()

	item, ok := c.items[key]
	if !ok {
		err = ErrNotFound
		return
	}

	if item.Expiration > 0 && item.Expiration < time.Now().UnixNano() {
		err = ErrNotFound
		return
	}

	return item.Value, nil
}

func (c *Cache) Delete(key string) (err error) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.items[key]; !ok {
		return ErrNotFound
	}

	delete(c.items, key)
	return
}

func (c *Cache) startGC() {
	go func() {
	L:
		for {
			select {
			case <-time.After(c.cleanupInterval):
				c.cleanExpiredKeys()
			case <-c.done:
				break L
			}
		}
		fmt.Println("memcache gracefull stop done!")
	}()
}

func (c *Cache) cleanExpiredKeys() {
	for key, item := range c.items {
		if item.Expiration > 0 && item.Expiration < time.Now().UnixNano() {
			delete(c.items, key)
		}
	}
}

func (c *Cache) Stop() (err error) {
	c.done <- struct{}{}
	return nil
}
