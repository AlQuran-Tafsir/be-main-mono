package repository

import (
	"context"
	"github.com/allegro/bigcache/v3"
	"github.com/alqurantafsir/be-main-monolith/domain"
	"os"
	"time"
)

type bigCacheRepository struct {
	bigCache *bigcache.BigCache
}

func NewBigCache() domain.CacheRepository {
	cache, err := bigcache.New(context.Background(), bigcache.DefaultConfig(24*30*12*time.Hour))
	if err != nil {
		os.Exit(1)
	}

	return &bigCacheRepository{
		bigCache: cache,
	}
}

func (b bigCacheRepository) Set(key string, entry []byte, d time.Duration) error {
	return b.bigCache.Set(key, entry)
}

func (b bigCacheRepository) Get(key string) ([]byte, error) {
	return b.bigCache.Get(key)
}

func (b bigCacheRepository) Delete(key string) error {
	return b.bigCache.Delete(key)
}
