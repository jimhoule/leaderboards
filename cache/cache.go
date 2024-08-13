package cache

type CacheHandler[TCache any] struct {
	Cache *TCache
}

type Cacher[TCache any] interface {
	New(address string, password string) *CacheHandler[TCache]
}

func NewCacheHandler[TCache any](cacher Cacher[TCache], address string, password string) *CacheHandler[TCache] {
	return cacher.New(address, password)
}