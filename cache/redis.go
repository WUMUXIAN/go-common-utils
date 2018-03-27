package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisCacher is an redis implementation of cacher.
type RedisCacher struct {
	p *redis.Pool
}

// Redis is the only one redis cacher, we don't support multiple intances of redis cacher for now.
var Redis *RedisCacher

// NewRedisCacher creates a new redis cacher
func NewRedisCacher(server, password string) error {
	// Has an existing pool, close it.
	if Redis != nil {
		Redis.ClosePool()
	}

	Redis = new(RedisCacher)
	if Redis.p == nil {
		Redis.p = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server)
				if err != nil {
					return nil, err
				}
				if password != "" {
					if _, err := c.Do("AUTH", password); err != nil {
						c.Close()
						return nil, err
					}
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("PING")
				return err
			},
		}
	}
	return Redis.p.Get().Err()
}

// ClosePool closes the redis pool
func (o *RedisCacher) ClosePool() {
	if o.p != nil {
		o.p.Close()
	}
}

// Set a key value pair, the value can be string, int64 and etc.
func (o *RedisCacher) Set(key string, value interface{}, expiration ...interface{}) error {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	_, err := redisConnection.Do("SET", key, value)

	if err == nil && expiration != nil {
		_, err = redisConnection.Do("EXPIRE", key, expiration[0])
	}
	if err != nil {
		fmt.Println("Set Cache Error", err)
	}
	return err
}

// Expire sets a expiration time for key
func (o *RedisCacher) Expire(key string, expiration int) error {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	_, err := redisConnection.Do("EXPIRE", key, expiration)
	return err
}

// TTL gets the remaining seconds for key
func (o *RedisCacher) TTL(key string) (int, error) {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	return redis.Int(redisConnection.Do("TTL", key))
}

// SetGob sets a key value pair, value will be gob encoded
func (o *RedisCacher) SetGob(key string, value interface{}, expiration ...interface{}) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(map[interface{}]interface{}{"value": value})

	if err == nil {
		err = o.Set(key, buffer.Bytes(), expiration...)
	}

	return err
}

// GetGob gets a gob encoded value from key
func (o *RedisCacher) GetGob(key string) (interface{}, error) {
	b, err := o.GetBytes(key)
	value := make(map[interface{}]interface{})
	if err == nil {
		buffer := bytes.NewBuffer(b)
		dec := gob.NewDecoder(buffer)
		err = dec.Decode(&value)
	}
	return value["value"], err
}

// GetString gets a string value from key
func (o *RedisCacher) GetString(key string) (string, error) {
	return redis.String(o.Get(key))
}

// GetBytes gets a []byte value from key
func (o *RedisCacher) GetBytes(key string) ([]byte, error) {
	return redis.Bytes(o.Get(key))
}

// GetInt64 gets a int64 value from key
func (o *RedisCacher) GetInt64(key string) (int64, error) {
	return redis.Int64(o.Get(key))
}

// GetFloat64 gets a float64 value from key
func (o *RedisCacher) GetFloat64(key string) (float64, error) {
	return redis.Float64(o.Get(key))
}

// Get a value from key
func (o *RedisCacher) Get(key string) (interface{}, error) {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	return redisConnection.Do("GET", key)
}

// Del a key
func (o *RedisCacher) Del(key string) {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	redisConnection.Do("DEL", key)
}

// Scan through the keys with given cursor, pattern and count
func (o *RedisCacher) Scan(cursor int, count int, pattern string) (nextCursor int, keys []string, err error) {
	redisConnection := o.p.Get()
	defer redisConnection.Close()
	result, err := redis.Values(redisConnection.Do("SCAN", cursor, "MATCH", pattern, "COUNT", count))
	if err != nil {
		return
	}
	nextCursor, err = redis.Int(result[0], nil)
	if err != nil {
		return
	}
	keys, err = redis.Strings(result[1], nil)
	if err != nil {
		return
	}
	return
}
