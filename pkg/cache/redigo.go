package cache

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	onceNewRedigoClient    sync.Once
	onceNewRedigoClientRes *RedigoClient
	onceNewRedigoClientErr error
)

// RedigoClient returns a redis client using redigo library.
type RedigoClient struct {
	pool *redis.Pool
}

// NewRedigoClient return a redis client.
func NewRedigoClient() (*RedigoClient, error) {
	onceNewRedigoClient.Do(func() {

		// Create connection pool.
		connPool := &redis.Pool{
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "127.0.0.1:8000")
			},
			DialContext: nil,
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Second {
					return nil
				}
				_, err := c.Do("PING")
				return err
			},
			MaxIdle:         10,
			MaxActive:       100,
			IdleTimeout:     0,
			Wait:            true,
			MaxConnLifetime: 0,
		}

		// Try to dial the redis.
		// On error close previous open connection pool.
		if _, err := connPool.Dial(); err != nil {
			_ = connPool.Close()
			onceNewRedigoClientErr = err
			return
		}

		onceNewRedigoClientRes = &RedigoClient{
			pool: connPool,
		}
	})

	return onceNewRedigoClientRes, onceNewRedigoClientErr
}

// Get gets the value from redis in []byte form.
func (r *RedigoClient) Get(ctx context.Context, key string) ([]byte, error) {
	const commandName = "GET"

	con := r.pool.Get()
	defer func() {
		_ = con.Close()
	}()

	data, err := redis.Bytes(con.Do(commandName, key))
	if err != nil && err != redis.ErrNil {
		return data, err
	}

	return data, nil
}

// SetEX sets the value to a key with timeout in seconds.
func (r *RedigoClient) SetEX(ctx context.Context, key string, seconds int64, value string) error {
	const commandName = "SET"

	con := r.pool.Get()
	defer func() {
		_ = con.Close()
	}()

	data, err := redis.String(con.Do(commandName, key, value, "ex", seconds))
	if err != nil && err != redis.ErrNil {
		return err
	}

	// Extra check for set operations.
	if err == redis.ErrNil || !strings.EqualFold("OK", data) {
		return err
	}

	return nil
}

// Exists checks whether the key exists in redis.
func (r *RedigoClient) Exists(ctx context.Context, key string) (bool, error) {
	const commandName = "EXISTS"

	con := r.pool.Get()
	defer func() {
		_ = con.Close()
	}()

	data, err := redis.Int64(con.Do(commandName, key))
	if err != nil {
		return false, err
	}
	if data != 1 {
		return false, nil
	}

	return true, nil
}
