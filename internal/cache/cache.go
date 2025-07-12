package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/dunamismax/go-modern-scaffold/internal/config"
)

// Cache holds the cache clients.
type Cache struct {
	Memory *ristretto.Cache
}

// New creates a new Cache instance with a Ristretto cache.
func New(cfg *config.Cache) (*Cache, error) {
	memCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: cfg.NumCounters,
		MaxCost:     cfg.MaxCost,
		BufferItems: cfg.BufferItems,
	})
	if err != nil {
		return nil, err
	}

	return &Cache{Memory: memCache}, nil
}

// Set adds an item to the memory cache.
func (c *Cache) Set(key string, value interface{}, cost int64) bool {
	return c.Memory.Set(key, value, cost)
}

// SetWithTTL adds an item to the memory cache with a TTL.
func (c *Cache) SetWithTTL(key string, value interface{}, cost int64, ttl time.Duration) bool {
	return c.Memory.SetWithTTL(key, value, cost, ttl)
}

// Get retrieves an item from the memory cache.
func (c *Cache) Get(key string) (interface{}, bool) {
	return c.Memory.Get(key)
}

// Del removes an item from the memory cache.
func (c *Cache) Del(key string) {
	c.Memory.Del(key)
}
