package cache

import (
	"sync"
	"time"

	"github.com/cyberhorsey/fliphousechill/marketcaps"
)

type CacheItem struct {
	Item     marketcaps.MarketCapResp `json:"item"`
	CachedAt time.Time                `json:"cached_at"`
}

type Cache struct {
	items map[string]CacheItem
	mutex sync.Mutex
}

func New() *Cache {
	return &Cache{
		items: make(map[string]CacheItem),
		mutex: sync.Mutex{},
	}
}

func (c *Cache) Get(key string) (marketcaps.MarketCapResp, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	item, found := c.items[key]
	if !found || time.Since(item.CachedAt) > marketcaps.MarketCaps[0].RefreshRate {
		return marketcaps.MarketCapResp{}, false
	}
	return item.Item, true
}

// in cache/cache.go
func (c *Cache) GetRaw(key string) (CacheItem, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	item, found := c.items[key]
	return item, found
}

func (c *Cache) Set(key string, item marketcaps.MarketCapResp) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = CacheItem{
		Item:     item,
		CachedAt: time.Now().UTC(),
	}
}

func (c *Cache) All() []CacheItem {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	copy := make([]CacheItem, 0, len(c.items))
	for _, v := range c.items {
		copy = append(copy, CacheItem{
			Item:     v.Item,
			CachedAt: v.CachedAt,
		})
	}
	return copy
}
