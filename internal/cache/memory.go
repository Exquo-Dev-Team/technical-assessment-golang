package cache

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