package cache

import "encoding/gob"

// Cacher defines the interface for all typs of cacher. e.g. redis, memcached and etc.
type Cacher interface {
	Set(key string, value interface{}, expiration ...interface{}) error
	Get(key string) (interface{}, error)
	Del(key string)
	Scan(cursor int, count int, pattern string) (nextCursor int, keys []string, err error)
	Expire(key string, expiration int) error
	TTL(key string) (int, error)
	SetGob(key string, value interface{}, expiration ...interface{})
	GetGob(key string) (interface{}, error)
	SetJSON(key string, value interface{}, expiration ...interface{})
	GetJSON(key string) (jsonBytes []byte, err error)
	HSet(hash, key string, value interface{}, expiration ...interface{}) error
	HGet(hash, key string) (interface{}, error)
	HINCRBY(hash, key string, value interface{}) error
}

// GobRegister registers models with gob.
func GobRegister(models ...interface{}) {
	for _, model := range models {
		gob.Register(model)
	}
}
