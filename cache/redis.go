package cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisCacher is an redis implementation of cacher.
type RedisCacher struct {
	p       *redis.Pool
	dbIndex int
}

// GetConn gets a connection
func (o *RedisCacher) GetConn() redis.Conn {
	conn := o.p.Get()
	if o.dbIndex != 0 {
		conn.Do("SELECT", o.dbIndex)
	}
	return conn
}

// Redis is the only one redis cacher, we don't support multiple intances of redis cacher for now.
var Redis *RedisCacher

// NewRedisCacher creates a new redis cacher
func NewRedisCacher(server, password string, dbIndex ...int) error {
	// Has an existing pool, close it.
	if Redis != nil {
		Redis.ClosePool()
	}

	Redis = new(RedisCacher)
	if len(dbIndex) > 0 {
		Redis.dbIndex = dbIndex[0]
	}

	if Redis.p == nil {
		Redis.p = &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", server, redis.DialConnectTimeout(time.Second*1))
				if err != nil {
					return nil, err
				}
				if password != "" {
					if _, err = c.Do("AUTH", password); err != nil {
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
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("SET", key, value)

	if err == nil && expiration != nil {
		_, err = redisConnection.Do("EXPIRE", key, expiration[0])
	}
	return err
}

// Expire sets a expiration time for key
func (o *RedisCacher) Expire(key string, expiration int) error {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("EXPIRE", key, expiration)
	return err
}

// TTL gets the remaining seconds for key
func (o *RedisCacher) TTL(key string) (int, error) {
	redisConnection := o.GetConn()
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

// SetJSON sets a key value pair, the value is a json
func (o *RedisCacher) SetJSON(key string, value interface{}, expiration ...interface{}) error {
	str, err := json.Marshal(value)
	if err == nil {
		err = o.Set(key, str, expiration...)
	}
	return err
}

// GetJSON gets a json value from key
func (o *RedisCacher) GetJSON(key string) (jsonBytes []byte, err error) {
	return o.GetBytes(key)
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
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	return redisConnection.Do("GET", key)
}

// Del a key
func (o *RedisCacher) Del(key string) {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	redisConnection.Do("DEL", key)
}

// Scan through the keys with given cursor, pattern and count
func (o *RedisCacher) Scan(cursor int, count int, pattern string) (nextCursor int, keys []string, err error) {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	result, err := redis.Values(redisConnection.Do("SCAN", cursor, "MATCH", pattern, "COUNT", count))
	if err != nil {
		return
	}
	nextCursor, _ = redis.Int(result[0], nil)
	keys, _ = redis.Strings(result[1], nil)

	return
}

// HSet sets a key:value in hash set.
func (o *RedisCacher) HSet(hash, key string, value interface{}, expiration ...interface{}) error {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("HSET", hash, key, value)

	if err == nil && expiration != nil {
		_, err = redisConnection.Do("EXPIRE", key, expiration[0])
	}
	return err
}

// HGet gets a value from hash set
func (o *RedisCacher) HGet(hash, key string) (interface{}, error) {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	return redisConnection.Do("HGET", hash, key)
}

// HINCRBY increments a value for hash set by key.
func (o *RedisCacher) HINCRBY(hash, key string, value interface{}) error {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("HINCRBY", hash, key, value)
	return err
}

// FlushAll flushes all keys in all db.
func (o *RedisCacher) FlushAll() error {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("FLUSHALL")
	return err
}

// Flush flushes all keys in selected db.
func (o *RedisCacher) Flush() error {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	_, err := redisConnection.Do("FLUSHDB")
	return err
}

// GetDBSize get the number of keys in the currently-selected database.
func (o *RedisCacher) GetDBSize() (interface{}, error) {
	redisConnection := o.GetConn()
	defer redisConnection.Close()
	return redisConnection.Do("DBSIZE")
}
