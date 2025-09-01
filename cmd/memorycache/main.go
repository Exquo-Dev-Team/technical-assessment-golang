package main

import (
	"errors"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, bool)
	Del(key string) bool
}

type MemoryCache struct {
}

func NewMemoryCache() *MemoryCache {
	return nil
}

func (c *MemoryCache) Set(key string, value any) error {
	return nil
}

func (c *MemoryCache) Get(key string) (any, bool) {
	return nil, false
}

func (c *MemoryCache) Del(key string) bool {
	return false
}

func main() {
	NewMemoryCache()
}
