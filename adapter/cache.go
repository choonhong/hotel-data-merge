package adapter

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/choonhong/hotel-data-merge/ent"
)

type Cache struct {
	*bigcache.BigCache
}

func NewCache() *Cache {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		panic(err)
	}

	return &Cache{cache}
}

func (c *Cache) Get(v interface{}) ([]*ent.Hotel, error) {
	key, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal params: %w", err)
	}

	data, err := c.BigCache.Get(string(key))
	if err != nil {
		return nil, err
	}

	hotels := []*ent.Hotel{}
	if err := json.Unmarshal(data, &hotels); err != nil {
		return nil, err
	}

	return hotels, nil
}

func (c *Cache) Set(v interface{}, hotels []*ent.Hotel) error {
	key, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("json.Marshal params: %w", err)
	}

	data, err := json.Marshal(hotels)
	if err != nil {
		return fmt.Errorf("json.Marshal hotels: %w", err)
	}

	return c.BigCache.Set(string(key), data)
}

func (c *Cache) Clear() error {
	return c.BigCache.Reset()
}

func (c *Cache) Close() {
	c.BigCache.Close()
}
