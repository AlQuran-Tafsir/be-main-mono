package domain

import "time"

type CacheRepository interface {
	Set(key string, entry []byte, d time.Duration) error
	Get(key string) ([]byte, error)
	Delete(key string) error
}
